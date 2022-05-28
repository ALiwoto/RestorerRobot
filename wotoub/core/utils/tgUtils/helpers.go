package tgUtils

import (
	wg "github.com/ALiwoto/wotoub/wotoub/core/wotoValues/wotoGlobals"
	"github.com/gotd/td/tg"
)

func GetIdFromPeerClass(p tg.PeerClass) int64 {
	switch v := p.(type) {
	case *tg.PeerUser: // peerUser#59511722
		return v.UserID
	case *tg.PeerChat: // peerChat#36c6019a
		return v.ChatID
	case *tg.PeerChannel: // peerChannel#a2a5371e
		return v.ChannelID
	default:
		return 0
	}
}

func GetInputUserFromId(id int64) *tg.InputUser {
	if GetUserFromIdHelper == nil {
		return nil
	}
	u, err := GetUserFromIdHelper(id)
	if u == nil || err != nil {
		return nil
	}

	return u
}

func GetInputUser(id, hash int64) *tg.InputUser {
	return &tg.InputUser{
		UserID:     id,
		AccessHash: hash,
	}
}

func GetInputUserFromUsername(username string) *tg.InputUser {
	contacts, err := wg.API.ContactsResolveUsername(gCtx, username)
	if err != nil || contacts == nil {
		return nil
	}

	if len(contacts.Users) > 0 {
		u, ok := contacts.Users[0].AsNotEmpty()
		if u != nil && ok {
			if SaveTgUser != nil {
				// try to cache the user
				go SaveTgUser(u)
			}

			return &tg.InputUser{
				UserID:     u.ID,
				AccessHash: u.AccessHash,
			}
		}
	}

	return nil

}
