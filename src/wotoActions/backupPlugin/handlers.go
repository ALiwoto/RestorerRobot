package backupPlugin

import (
	"os"

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

		err = backupUtils.ZipSource(sourceFileName, finalFileName)
		if err != nil {
			_, _ = container.ReplyText("Failed to zip backup file" + err.Error())
			return em.ErrEndGroups
		}
		_ = os.Remove(sourceFileName)

		container.UploadFileToChatsByPath(finalFileName, &em.UploadDocumentToChatsOptions{
			ChatIDs:    targetChats,
			Goroutines: 60,
			// Caption: "TODO",
		})
	}

	return em.ErrEndGroups
}
