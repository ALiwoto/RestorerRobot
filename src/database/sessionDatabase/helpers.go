package sessionDatabase

import (
	"strings"
	"sync"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/tgUtils"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues"
	wg "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/gotd/td/tg"
	"gorm.io/gorm"
)

func init() {
	tgUtils.GetUserFromIdHelper = GetInputUserFromId
	tgUtils.GetInputPeerInfo = GetPeerInfoFromId
	tgUtils.SaveTgUser = SaveTgUser
}

func StartDatabase(db *gorm.DB, mut *sync.Mutex) error {
	dbSession = db
	dbMutex = mut
	return nil
}

func SaveTgUser(u *tg.User) error {
	return SaveUser(u.ID, u.AccessHash)
}

func SaveTgChannel(ch *tg.Channel) error {
	// we will be adding -100 to the channel id
	idStr := ssg.ToBase10(ch.ID)
	if !strings.HasPrefix(idStr, "-100") {
		idStr = "-100" + idStr
	}

	return SaveChannel(ssg.ToInt64(idStr), ch.AccessHash)
}

func SaveChannel(id int64, hash int64) error {
	var u *wg.PeerInfo
	u = peerDbMap.Get(id)
	if u == nil || u.AccessHash != hash {
		if u == nil {
			u = &wg.PeerInfo{
				PeerId:     id,
				AccessHash: hash,
				PeerType:   wg.PeerTypeChannel,
			}
		}

		return NewPeerInfo(u)
	}
	return nil
}

func SaveUser(id int64, hash int64) error {
	var u *wg.PeerInfo
	u = peerDbMap.Get(id)
	if u == nil || u.AccessHash != hash {
		if u == nil {
			u = &wg.PeerInfo{
				PeerId:     id,
				AccessHash: hash,
				PeerType:   wg.PeerTypeUser,
			}
		}

		return NewPeerInfo(u)
	}
	return nil
}

func NewPeerInfo(u *wg.PeerInfo) error {
	dbMutex.Lock()
	tx := dbSession.Begin()
	tx.Save(u)
	tx.Commit()
	dbMutex.Unlock()
	peerDbMap.Add(u.PeerId, u)
	return nil
}

func GetPeerInfoFromId(id int64) (*wg.PeerInfo, error) {
	if dbSession == nil {
		return nil, ErrNoSession
	}

	u := peerDbMap.Get(id)
	if u != nil {
		return u, nil
	}

	u = &wg.PeerInfo{}
	dbMutex.Lock()
	dbSession.Model(ModelUser).Where("peer_id = ?", id).Take(u)
	dbMutex.Unlock()
	if u.PeerId != id || u.AccessHash == 0 {
		if wotoValues.IsRealOwner(id) {
			u.AccessHash = wg.Self.AccessHash
			u.PeerId = id
			u.PeerType = wg.PeerTypeUser
			err := NewPeerInfo(u)
			if err != nil {
				return nil, err
			}
			return u, nil
		}
		// not found
		return nil, nil
	}

	peerDbMap.Add(u.PeerId, u)

	return u, nil
}

func GetInputUserFromId(id int64) (*tg.InputUser, error) {
	u, err := GetPeerInfoFromId(id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, nil
	}
	return &tg.InputUser{
		UserID:     u.PeerId,
		AccessHash: u.AccessHash,
	}, nil
}
