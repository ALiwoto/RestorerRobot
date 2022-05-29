package statsPlugin

import (
	"log"

	em "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoStyle"
)

func wotoUbHandler(container *em.WotoContainer) error {
	message := container.Message
	userId := container.GetEffectiveUserID()
	if userId != 1341091260 {
		return nil
	}
	txt := wotoStyle.GetBold("wotoub ")
	txt.AppendMonoThis("v1.0.1")
	txt.AppendBoldThis("\n • Status: ")
	txt.AppendMonoThis("Active")
	container.ReplyStyledText(txt)

	//wv.SenderHelper.Reply(*container.Entities, container.GetAnswerable()).Text(container.Ctx(), "hello")
	//sender.Resolve("1117157532").Text(wv.GCtx, "Hello")
	log.Println(message.Message)
	return em.ErrEndGroups
}
