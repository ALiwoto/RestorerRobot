package funPlugin

import (
	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/tgUtils"
	em "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoStyle"
)

func mentionHandler(container *em.WotoContainer) error {
	message := container.Message
	args := container.Args()
	if args == nil {
		return em.ErrEndGroups
	}

	var usernameTarget string
	target, ok := args.GetAsIntegerOrRaw("u", "user", "target")
	if target == 0 || !ok {
		usernameTarget = args.GetAsStringOrRaw("u", "user", "target")
		if usernameTarget == "" {
			usernameTarget = args.GetFirstNoneEmptyValue()
		}
	}

	txt := wotoStyle.GetBold("User mention: ")

	if usernameTarget == "" {
		if target == 0 {
			target = tgUtils.GetIdFromPeerClass(message.FromID)
		}

		txt.AppendMentionThis("here", target)
	} else {
		txt.AppendMentionUsernameThis("here", usernameTarget)
	}

	_, _ = container.ReplyStyledText(txt)
	return em.ErrEndGroups
}

func fishHandler(container *em.WotoContainer) error {
	txt := wotoStyle.GetBold("/fish")

	_, _ = container.ReplyStyledText(txt)
	return em.ErrEndGroups
}
