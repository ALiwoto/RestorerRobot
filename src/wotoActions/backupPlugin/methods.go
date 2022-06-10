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
	time.Sleep(time.Minute)
	m.checkBackups() // first run is necessary

	for {
		time.Sleep(m.checkInterval)
		m.checkBackups()
	}
}

func (m *BackupScheduleManager) checkBackups() {
	var currentContainers []*BackupScheduleContainer
	for i := 0; i < len(m.containers); i++ {
		if m.containers[i].IsWithin(managerTimeInterval) && m.containers[i].currentInfo != nil {
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

func (m *BackupScheduleManager) PrepareBackupInfo(currentContainers []*BackupScheduleContainer) {
	md := wotoStyle.GetBold("🔹 Following databases will be backed up within the next " + managerTimeInterval.String() + ":")

	var current *BackupScheduleContainer
	for i := 0; i < len(currentContainers); i++ {
		current = currentContainers[i]
		current.currentInfo = backupDatabase.GenerateBackupInfo(current.RemainingTime(), 0)
		md.ElThis().AppendThis(current.ParseAsMd())
	}

	sender := message.NewSender(wotoGlobals.API)

	for _, chatId := range m.ChatIDs {
		inputTarget, err := tgUtils.GetInputPeerClass(chatId)
		if err != nil {
			return
		}

		target := sender.To(inputTarget)
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
	return time.Since(m.LastBackupDate)-m.BackupInterval >= d
}

func (m *BackupScheduleContainer) RemainingTime() time.Duration {
	tmp := m.BackupInterval - time.Since(m.LastBackupDate)
	if tmp <= 0 {
		return 0
	}

	return tmp
}

func (c *BackupScheduleContainer) ParseAsMd() wotoStyle.WStyle {
	return wotoStyle.GetHyperLink(c.DatabaseConfig.GetSectionName(), "http://t.me/ShigeoRobot?start=ghelp_sibylsystem")
}

func (c *BackupScheduleContainer) RunBackup() {
	if c.currentInfo == nil {
		return
	}

	time.Sleep(c.RemainingTime())

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

	err = backupUtils.BackupDatabase(theUrl, sourceFileName, bType)
	if err != nil {
		//_, _ = container.ReplyError("Failed to backup database", err)
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
		BackupInitType: "Automatic Backup System",
		InitiatedBy:    "Backup Schedule Container",
		UserId:         0,
		DateTime:       time.Now(),
		FileSize:       sourceFileSize,
		BackupFormat:   strings.ToUpper(bType),
	}

	err = backupUtils.ZipSource(sourceFileName, finalFileName)
	if err != nil {
		// _, _ = container.ReplyError("Failed to zip backup file", err)
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
		// _, _ = container.ReplyError("Failed to upload backup file", err)
		return
	}

	c.currentInfo = nil
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
