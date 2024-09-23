package statsPlugin

import (
	"github.com/ALiwoto/RestorerRobot/src/core/wotoConfig"
	em "github.com/ALiwoto/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/ALiwoto/RestorerRobot/src/core/wotoStyle"
	"github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
)

func statusHandler(container *em.WotoContainer) error {
	// message := container.Message
	userId := container.GetEffectiveUserID()
	if !wotoConfig.IsOwner(userId) {
		return em.ErrEndGroups
	}

	// input := tgUtils.GetInputUserFromId(userId)
	// print(input)

	txt := wotoStyle.GetEmpty()
	txt.Bold("🔹 @" + wotoGlobals.Self.Username + "'s stats:")
	txt.Bold("\n • Version: ").Mono(wotoGlobals.AppVersion)
	txt.Bold("\n • Status: ")
	txt.Mono("Active")
	fetchGitStats(txt)
	container.ReplyStyledText(txt)

	// err := container.UploadFileToChatByPath("spy.mkv", -1001695105982, wotoStyle.GetBold("hi"))
	// err := container.UploadFileToChatByPath("run.sh", -1001695105982, wotoStyle.GetBold("hi"))
	// if err != nil {
	// logging.Error(err)
	// }

	//wv.SenderHelper.Reply(*container.Entities, container.GetAnswerable()).Text(container.Ctx(), "hello")
	//sender.Resolve("1117157532").Text(wv.GCtx, "Hello")
	// log.Println(message.Message)
	return em.ErrEndGroups
}
