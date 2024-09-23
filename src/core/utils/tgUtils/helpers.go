package tgUtils

import (
	wg "github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
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

func GetInputPeerClass(id int64) (tg.InputPeerClass, error) {
	info, err := GetInputPeerInfo(id)
	if err != nil {
		return nil, err
	}

	if info == nil {
		return nil, wg.ErrPeerIdInvalid
	}

	switch info.PeerType {
	case wg.PeerTypeEmpty:
		return &tg.InputPeerEmpty{}, nil
	case wg.PeerTypeSelf:
		return &tg.InputPeerSelf{}, nil
	case wg.PeerTypeChat:
		return &tg.InputPeerChat{
			ChatID: info.GetRealID(),
		}, nil
	case wg.PeerTypeChannel:
		return &tg.InputPeerChannel{
			ChannelID:  info.GetRealID(),
			AccessHash: info.AccessHash,
		}, nil
	case wg.PeerTypeUser:
		return &tg.InputPeerUser{
			UserID:     info.GetRealID(),
			AccessHash: info.AccessHash,
		}, nil
	case wg.PeerTypeUserFromMessage:
		return &tg.InputPeerUserFromMessage{
			UserID: info.GetRealID(),
		}, nil
	}
	return nil, nil
}

func GetUserIdFromPeerClass(p tg.PeerClass) int64 {
	peerUser, ok := p.(*tg.PeerUser)
	if ok {
		return peerUser.UserID
	}
	return 0
}

func GetEffectiveUserIdFromMessage(msg *tg.Message) int64 {
	switch {
	case msg.FromID != nil:
		return GetUserIdFromPeerClass(msg.FromID)
	case msg.PeerID != nil:
		return GetUserIdFromPeerClass(msg.PeerID)
	}

	return 0
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
