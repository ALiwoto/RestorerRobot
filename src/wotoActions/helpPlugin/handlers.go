package helpPlugin

import (
	"strings"
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	em "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoStyle"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/AnimeKaizoku/RestorerRobot/src/database/backupDatabase"
	"github.com/AnimeKaizoku/RestorerRobot/src/wotoActions/backupPlugin"
	"github.com/AnimeKaizoku/ssg/ssg"
)

func configsHandler(container *em.WotoContainer) error {
	userId := container.GetEffectiveUserID()
	user := container.Entities.Users[userId]
	if user == nil || !wotoConfig.IsOwner(userId) {
		return em.ErrEndGroups
	}

	var currentName string
	var currentLastBackupDate time.Time

	if len(wotoConfig.WotoConf.Sections) == 0 {
		_, _ = container.ReplyText("No databases configured yet.")
		return em.ErrEndGroups
	}

	md := wotoStyle.GetBold("📄 List of configured databases:")
	for i, current := range wotoConfig.WotoConf.Sections {
		currentName = current.GetSectionName()
		md.Bold("\n" + ssg.ToBase10(int64(i)) + "- ").Mono(currentName)
		currentLastBackupDate = backupDatabase.GetLastBackupDate(currentName)
		if !currentLastBackupDate.IsZero() {
			md.Bold("[").Mono(currentLastBackupDate.Format("Mon-01-02-2006 03:04 PM")).Bold("]")
		}
	}

	_, _ = container.ReplyStyledText(md)

	return em.ErrEndGroups
}

func startHandler(container *em.WotoContainer) error {
	userId := container.GetEffectiveUserID()
	user := container.Entities.Users[userId]
	if user == nil {
		return em.ErrEndGroups
	}

	if !wotoConfig.IsOwner(userId) {
		md := wotoStyle.GetNormal("Hello there ").MentionHash(user.FirstName, user.ID, user.AccessHash)
		md.Normal("! This bot is for internal usage only, please contact us through ")
		md.Link("support group", "https://t.me/KaizokuBots")
		md.Normal(" if you need any help regarding this matter.")
		_, _ = container.ReplyStyledText(md)
		return em.ErrEndGroups
	}

	text := container.Message.Message

	// fast way of checking to see if this command has any arguments
	// or not.
	if len(text) > 5 && strings.Contains(text, " ") {
		myStrs := strings.Split(text, " ")
		switch {
		case wotoGlobals.IsBackupUniqueId(myStrs[1]):
			bInfo := backupDatabase.GetBackupInfo(wotoGlobals.BackupUniqueIdValue(myStrs[1]))
			if bInfo == nil {
				_, _ = container.ReplyText("No backup info found with this ID.")
				return em.ErrEndGroups
			}

			md := wotoStyle.GetBold("🔹 Backup info:")
			md.Bold("\n・Name: ").Mono(bInfo.DatabaseName)
			md.Bold("\n・Unique ID: ").Mono(string(bInfo.BackupUniqueId))
			md.Bold("\n・Last backup: ").Mono(bInfo.BackupDate.Format("2006-01-02 15:04:05"))
			md.Bold("\n・Backup status: ").Mono(bInfo.GetStrStatus())
			if bInfo.Message != "" {
				switch bInfo.Status {
				case wotoGlobals.BackupStatusFailed:
					md.Bold("\n・Error message: ").Mono(bInfo.Message)
				default:
					md.Bold("\n・Notice: ").Mono(bInfo.Message)
				}
			}
			_, _ = container.ReplyStyledText(md)
			return em.ErrEndGroups
		}
	}

	md := wotoStyle.GetNormal("Welcome to " + wotoGlobals.Self.Username)
	md.Normal(" master!")
	md.Normal("\n This bot lets you take backup from unlimited amount of databases ")
	md.Normal("with the specified time intervals.")
	_, _ = container.ReplyStyledText(md)
	return em.ErrEndGroups
}

func dbInfoHandler(container *em.WotoContainer) error {
	userId := container.GetEffectiveUserID()
	user := container.Entities.Users[userId]
	if user == nil || !wotoConfig.IsOwner(userId) {
		return em.ErrEndGroups
	}

	myStrs := ssg.SplitN(container.GetMessageText(), 2, " ", "\n")
	if len(myStrs) < 2 {
		_, _ = container.ReplyText("Please provide me a database name to get info of it.")
		return em.ErrEndGroups
	}

	dbName := myStrs[1]

	section := wotoConfig.GetSectionValueByName(dbName)
	if section == nil {
		_, _ = container.ReplyText("No such config section: " + dbName)
		return em.ErrEndGroups
	}

	md := wotoStyle.GetBold("🔸Database Info")
	md.Bold("\n・Name: ").Mono(dbName)
	md.Bold("\n・Type: ").Mono(section.BackupType)

	backupContainer := backupPlugin.GetContainerByName(dbName)
	if backupContainer != nil {
		md.Bold("\n・Time to next backup: ").Mono(backupContainer.RemainingTimeStr())
	}

	_, _ = container.ReplyStyledText(md)

	return em.ErrEndGroups
}
