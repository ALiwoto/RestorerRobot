package statsPlugin

import (
	"log"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/tgUtils"
	em "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoStyle"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
)

func statusHandler(container *em.WotoContainer) error {
	message := container.Message
	userId := container.GetEffectiveUserID()
	if userId != 1341091260 {
		return nil
	}

	input := tgUtils.GetInputUserFromId(userId)
	print(input)

	txt := wotoStyle.GetBold("@" + wotoGlobals.Self.Username + ":")
	txt.AppendMonoThis("v1.0.1")
	txt.AppendBoldThis("\n • Status: ")
	txt.AppendMonoThis("Active")
	container.ReplyStyledText(txt)
	err := container.UploadFileToChatByPath("run.sh", -1001695105982, wotoStyle.GetBold("hi"))
	print(err)

	//wv.SenderHelper.Reply(*container.Entities, container.GetAnswerable()).Text(container.Ctx(), "hello")
	//sender.Resolve("1117157532").Text(wv.GCtx, "Hello")
	log.Println(message.Message)
	return em.ErrEndGroups
}
