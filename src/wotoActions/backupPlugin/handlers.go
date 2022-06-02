package backupPlugin

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/backupUtils"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	em "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
)

func forceBackupHandler(container *em.WotoContainer) error {
	// message := container.Message
	userId := container.GetEffectiveUserID()
	if !wotoConfig.IsOwner(userId) {
		return em.ErrContinueGroups
	}

	userInfo, ok := container.Entities.Users[userId]
	if !ok {
		return em.ErrContinueGroups
	}

	theName := userInfo.FirstName + " " + userInfo.LastName
	if len(theName) > 32 {
		theName = theName[:32]
	}

	args, err := container.GetArgs()
	if err != nil {
		_, _ = container.ReplyError("Failed to parse arg", err)
		return em.ErrEndGroups
	}

	name := args.GetAsStringOrRaw("name", "url", "db", "database")
	isPrivate := args.GetAsBool("private", "isPrivate", "is-private")
	bType := args.GetAsString("type", "b-type", "bType", "backup-type", "backupType")
	if name == "" {
		name = args.GetFirstNoneEmptyValue()
	}

	if name == "" {
		_, _ = container.ReplyText("Please specify a name or a database url.")
		return em.ErrEndGroups
	}

	if bType == "" {
		bType = wotoConfig.BackupTypeDump // default is .dump
	}

	isUrl := wotoGlobals.IsDatabaseUrl(name)
	var theUrl string         // the url of the database we have to pass to backup helper function
	var targetChats []int64   // the chats we want to send our files to
	var sourceFileName string // the uncompressed backup file (output of the backup command)
	var originFileName string // the origin name that we have to append extensions to it
	var finalFileName string  // the file to be uploaded to tg
	var sourceFileSize string // the file size in this format: 10MB or 10.5MB

	if !isPrivate {
		targetChats = append(targetChats, wotoConfig.GetGlobalLogChannels()...)
	}
	if !isUrl {
		section := wotoConfig.GetSectionValueByName(name)
		if section == nil {
			_, _ = container.ReplyText("No such config section: " + name)
			return em.ErrEndGroups
		} else if !wotoGlobals.IsDatabaseUrl(section.DbUrl) {
			_, _ = container.ReplyText("You have provided wrong url format for the section " + name)
			return em.ErrEndGroups
		}

		theUrl = section.DbUrl
		sectionName := section.GetSectionName()
		originFileName = wotoConfig.GetBaseDirForBackup(sectionName) +
			backupUtils.GenerateFileNameFromValue(sectionName)
		sourceFileName = originFileName + "." + bType
		finalFileName = originFileName + wotoConfig.CompressedFileExtension
		targetChats = append(targetChats, section.LogChannels...)
		targetChats = append(targetChats, userId)

		err = backupUtils.BackupDatabase(theUrl, sourceFileName, bType)
		if err != nil {
			_, _ = container.ReplyText("Failed to backup database" + err.Error())
			return em.ErrEndGroups
		}

		// fetch file size
		fileInfo, err := os.Stat(sourceFileName)
		if err == nil {
			// format the file size
			sourceFileSize = backupUtils.FormatFileSize(fileInfo.Size())
		}

		captionOptions := &backupUtils.GenerateCaptionOptions{
			ConfigName:     sectionName,
			BackupInitType: "Manual Backup",
			InitiatedBy:    theName,
			UserId:         userId,
			DateTime:       time.Now(),
			FileSize:       sourceFileSize,
			BackupFormat:   strings.ToUpper(bType),
		}

		err = backupUtils.ZipSource(sourceFileName, finalFileName)
		if err != nil {
			_, _ = container.ReplyText("Failed to zip backup file" + err.Error())
			return em.ErrEndGroups
		}
		_ = os.Remove(sourceFileName)

		err = container.UploadFileToChatsByPath(finalFileName, &em.UploadDocumentToChatsOptions{
			ChatIDs:    targetChats,
			Goroutines: 60,
			Caption:    backupUtils.GenerateCaption(captionOptions),
		})
		if err != nil {
			_, _ = container.ReplyText("Failed to upload backup file" + err.Error())
			return em.ErrEndGroups
		}

		return em.ErrEndGroups
	}

	// if we are here, then we have a url
	theUrl = name
	sectionName := filepath.Base(theUrl) // dummy sectionName
	originFileName = wotoConfig.GetBaseDirForBackup(sectionName) +
		backupUtils.GenerateFileNameFromValue(sectionName)
	sourceFileName = originFileName + "." + bType
	finalFileName = originFileName + wotoConfig.CompressedFileExtension
	targetChats = append(targetChats, userId)
	err = backupUtils.BackupDatabase(theUrl, sourceFileName, bType)
	if err != nil {
		_, _ = container.ReplyError("Failed to backup database", err)
		return em.ErrEndGroups
	}

	captionOptions := &backupUtils.GenerateCaptionOptions{
		ConfigName:     sectionName,
		BackupInitType: "Manual Backup",
		InitiatedBy:    theName,
		UserId:         userId,
		DateTime:       time.Now(),
		FileSize:       sourceFileSize,
		BackupFormat:   strings.ToUpper(bType),
	}

	err = backupUtils.ZipSource(sourceFileName, finalFileName)
	if err != nil {
		_, _ = container.ReplyError("Failed to zip backup file", err)
		return em.ErrEndGroups
	}
	_ = os.Remove(sourceFileName)

	err = container.UploadFileToChatsByPath(finalFileName, &em.UploadDocumentToChatsOptions{
		FileName:   filepath.Base(finalFileName),
		ChatIDs:    targetChats,
		Goroutines: 60,
		Caption:    backupUtils.GenerateCaption(captionOptions),
	})
	if err != nil {
		_, _ = container.ReplyText("Failed to upload backup file" + err.Error())
		return em.ErrEndGroups
	}

	return em.ErrEndGroups
}
