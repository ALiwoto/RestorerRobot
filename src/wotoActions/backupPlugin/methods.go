package backupPlugin

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/backupUtils"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/logging"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/tgUtils"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	em "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoStyle"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/AnimeKaizoku/RestorerRobot/src/database/backupDatabase"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/uploader"
)

// --------------------------------------------------------

func (m *BackupScheduleManager) RunChecking() {
	time.Sleep(20 * time.Second)
	m.checkBackups() // first run is necessary

	for {
		time.Sleep(m.checkInterval)
		m.checkBackups()
	}
}

func (m *BackupScheduleManager) checkBackups() {
	var current *BackupScheduleContainer
	var currentContainers []*BackupScheduleContainer
	for i := 0; i < len(m.containers); i++ {
		current = m.containers[i]
		if current.NeedsImmediateBackup() ||
			(current.IsWithin(managerTimeInterval) && current.currentInfo == nil) {
			currentContainers = append(currentContainers, m.containers[i])
		}
	}

	if len(currentContainers) == 0 {
		return
	}

	m.PrepareBackupInfo(currentContainers)

	for i := 0; i < len(currentContainers); i++ {
		go currentContainers[i].RunBackup()
	}
}

// FormatInterval method will convert the time interval with this format:
// "1 day, 2 hours, 3 minutes, 4 seconds" (it will omit the smaller units if they are 0)
func (m *BackupScheduleManager) FormatInterval() string {
	myStr := managerTimeInterval.String()
	if strings.HasSuffix(myStr, "0s") {
		myStr = strings.TrimSuffix(myStr, "0s")
	}

	if strings.HasSuffix(myStr, "0m") {
		myStr = strings.TrimSuffix(myStr, "0m")
	}

	return strings.ReplaceAll(myStr, "h", "hours")
}

func (m *BackupScheduleManager) PrepareBackupInfo(currentContainers []*BackupScheduleContainer) {
	md := wotoStyle.GetBold("🔹 Following databases will be backed up in less than " + m.FormatInterval() + ":")

	var current *BackupScheduleContainer
	for i := 0; i < len(currentContainers); i++ {
		current = currentContainers[i]
		current.currentInfo = backupDatabase.GenerateBackupInfo(
			current.GetName(),
			current.RemainingTime(),
			0,
		)
		md.ElThis().AppendThis(current.ParseAsMd())
	}

	sender := message.NewSender(wotoGlobals.API)

	for _, chatId := range m.ChatIDs {
		inputTarget, err := tgUtils.GetInputPeerClass(chatId)
		if err != nil {
			return
		}

		target := sender.To(inputTarget).NoWebpage()
		if _, err := target.StyledText(context.Background(), md.GetStylingArray()...); err != nil {
			logging.Error(err)
		}
	}
}

func (m *BackupScheduleManager) convertToBackupInterval(days int) time.Duration {
	return time.Duration(days) * 24 * time.Hour
}

// --------------------------------------------------------

func (m *BackupScheduleContainer) IsWithin(d time.Duration) bool {
	m.mut.Lock()
	defer m.mut.Unlock()

	return !m.isSleeping && time.Since(m.LastBackupDate)-m.BackupInterval >= d
}

func (m *BackupScheduleContainer) RemainingTime() time.Duration {
	tmp := m.BackupInterval - time.Since(m.LastBackupDate)
	if tmp <= 0 {
		return 0
	}

	return tmp
}

func (c *BackupScheduleContainer) ParseAsMd() wotoStyle.WStyle {
	name := c.DatabaseConfig.GetSectionName()
	username := wotoGlobals.Self.Username
	link := "https://t.me/" + username + "?start=" + string(c.currentInfo.BackupUniqueId)
	return wotoStyle.GetHyperLink("・"+name, link)
}

func (c *BackupScheduleContainer) GetStrUniqueId() string {
	return string(c.currentInfo.BackupUniqueId)
}

func (c *BackupScheduleContainer) GetName() string {
	return c.DatabaseConfig.GetSectionName()
}

func (c *BackupScheduleContainer) GetBackupType() string {
	return c.DatabaseConfig.BackupType
}

func (c *BackupScheduleContainer) RunBackup() {
	c.mut.Lock()
	defer c.mut.Unlock()

	if c.currentInfo == nil || c.isSleeping {
		return
	}

	c.isSleeping = true
	time.Sleep(c.RemainingTime())

	setError := func(theErr error) {
		c.currentInfo.SetAsFailed(theErr)
		backupDatabase.UpdateBackupInfo(c.currentInfo)
		c.currentInfo = nil
	}

	section := c.DatabaseConfig
	var err error             // failed err reason
	var theUrl string         // the url of the database we have to pass to backup helper function
	var targetChats []int64   // the chats we want to send our files to
	var sourceFileName string // the uncompressed backup file (output of the backup command)
	var originFileName string // the origin name that we have to append extensions to it
	var finalFileName string  // the file to be uploaded to tg
	var sourceFileSize string // the file size in this format: 10MB or 10.5MB
	var bType string          // the backup type in string

	theUrl = section.DbUrl
	if section.BackupType != "" {
		bType = section.BackupType
	}

	sectionName := section.GetSectionName()
	originFileName = wotoConfig.GetBaseDirForBackup(sectionName) +
		backupUtils.GenerateFileNameFromValue(sectionName)
	sourceFileName = originFileName + "." + bType
	finalFileName = originFileName + wotoConfig.CompressedFileExtension
	targetChats = append(targetChats, section.LogChannels...)
	targetChats = append(targetChats, c.ChatIDs...)

	err = backupUtils.BackupDatabase(theUrl, sourceFileName, bType)
	if err != nil {
		setError(err)
		return
	}

	// fetch file size
	fileInfo, err := os.Stat(sourceFileName)
	if err == nil {
		// format the file size
		sourceFileSize = backupUtils.FormatFileSize(fileInfo.Size())
	}

	captionOptions := &backupUtils.GenerateCaptionOptions{
		ConfigName:     sectionName,
		BackupInitType: "Automatic Backup",
		InitiatedBy:    "Backup Schedule Container",
		UserId:         0,
		DateTime:       time.Now(),
		FileSize:       sourceFileSize,
		BackupFormat:   strings.ToUpper(bType),
	}

	err = backupUtils.ZipSource(sourceFileName, finalFileName)
	if err != nil {
		setError(err)
		return
	}
	_ = os.Remove(sourceFileName)

	err = c.UploadFileToChats(finalFileName, &em.UploadDocumentToChatsOptions{
		FileName:   filepath.Base(finalFileName),
		ChatIDs:    targetChats,
		Goroutines: 60,
		Caption:    backupUtils.GenerateCaption(captionOptions),
	})
	if err != nil {
		setError(err)
		return
	}

	dbInfo := backupDatabase.GetDatabaseInfo(c.DatabaseConfig.GetSectionName())
	dbInfo.LastBackup = time.Now()
	dbInfo.LastBackupUniqueId = c.currentInfo.BackupUniqueId
	backupDatabase.UpdateDatabaseInfo(dbInfo)

	c.currentInfo.SetAsFinished()
	backupDatabase.UpdateBackupInfo(c.currentInfo)

	c.LastBackupDate = dbInfo.LastBackup

	c.currentInfo = nil
}

func (c *BackupScheduleContainer) NeedsImmediateBackup() bool {
	// There are two conditions to trigger an immediate backup:
	// 1- the database has 0 finished backup.
	// 2- the database's last backup info is in pending status,
	//  which means the last time we have tried to backup the database,
	//  the program either got interrupted (such as being killed by user), or
	//  got stuck in a loop (such as the database is not responding).

	// case 1
	if backupDatabase.GetBackupFinishedCount(c.GetName()) == 0 {
		return true
	}

	// case 2
	lastStatus := backupDatabase.GetLastBackupStatus(c.GetName())
	return lastStatus.IsPending() || lastStatus.IsCanceled()
}

func (m *BackupScheduleContainer) UploadFileToChats(filename string, opts *em.UploadDocumentToChatsOptions) error {
	uploader := uploader.NewUploader(wotoGlobals.API).WithThreads(opts.Goroutines)
	sender := message.NewSender(wotoGlobals.API).WithUploader(uploader)
	upload, err := uploader.FromPath(context.Background(), filename)
	if err != nil {
		return err
	}
	if opts.Caption == nil {
		opts.Caption = wotoStyle.GetEmpty()
	}

	builder := message.UploadedDocument(upload, opts.Caption.GetStylingArray()...)
	if opts.FileName == "" {
		// user-provided file name is empty, fallback to using path.Base
		builder = builder.Filename(path.Base(filename))
	} else {
		builder.ForceFile(true).Filename(opts.FileName)
	}

	for _, chatID := range opts.ChatIDs {
		inputTarget, err := tgUtils.GetInputPeerClass(chatID)
		if err != nil {
			return err
		}

		target := sender.To(inputTarget)

		// Sending message with media.
		if _, err := target.Media(context.Background(), builder); err != nil {
			return err
		}
	}

	return nil
}

// --------------------------------------------------------
// --------------------------------------------------------
