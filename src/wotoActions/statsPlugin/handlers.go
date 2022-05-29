package statsPlugin

import (
	"log"

	em "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoStyle"
	"github.com/gotd/td/tg"
)

func wotoUbHandler(container *em.WotoContainer) error {
	message := container.Message
	from := message.FromID
	switch fromC := from.(type) {
	case *tg.PeerUser: // peerUser#59511722
		if fromC.UserID == 1341091260 {

			txt := wotoStyle.GetBold("wotoub ")
			txt.AppendMonoThis("v1.0.1")
			txt.AppendBoldThis("\n • Status: ")
			txt.AppendMonoThis("Active")
			container.ReplyStyledText(txt)

			//wv.SenderHelper.Reply(*container.Entities, container.GetAnswerable()).Text(container.Ctx(), "hello")
			//sender.Resolve("1117157532").Text(wv.GCtx, "Hello")
		}
	case *tg.PeerChat: // peerChat#36c6019a
	case *tg.PeerChannel: // peerChannel#a2a5371e
		//default:
	}
	log.Println(message.Message)
	return em.ErrEndGroups
}
