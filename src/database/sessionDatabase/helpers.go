package sessionDatabase

import (
	"strings"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/logging"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/tgUtils"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues"
	wg "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/gotd/td/tg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	tgUtils.GetUserFromIdHelper = GetInputUserFromId
	tgUtils.GetInputPeerInfo = GetPeerInfoFromId
	tgUtils.SaveTgUser = SaveTgUser
}

func StartDatabase() error {
	// check if `SESSION` variable is already established or not.
	// if yes, check if we have got any error from it or not.
	// if there is an error in the session, it mean we have to establish
	// a new connection again.
	if dbSession != nil && dbSession.Error == nil {
		return nil
	}

	var db *gorm.DB
	var err error
	var conf *gorm.Config
	if wotoConfig.IsDebug() {
		conf = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	} else {
		conf = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	db, err = gorm.Open(sqlite.Open(DbPath), conf)
	if err != nil {
		return err
	}

	dbSession = db

	logging.Info("Database connected")

	// Create tables if they don't exist
	err = dbSession.AutoMigrate(modelUser)
	if err != nil {
		return err
	}

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
	dbSession.Model(modelUser).Where("peer_id = ?", id).Take(u)
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
