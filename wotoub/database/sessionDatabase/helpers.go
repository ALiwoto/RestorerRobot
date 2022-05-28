package sessionDatabase

import (
	"github.com/ALiwoto/wotoub/wotoub/core/utils/logging"
	"github.com/ALiwoto/wotoub/wotoub/core/utils/tgUtils"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoConfig"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoValues"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoValues/wotoGlobals"
	"github.com/gotd/td/tg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	tgUtils.GetUserFromIdHelper = func(id int64) (*tg.InputUser, error) {
		u, err := GetUserFromId(id)
		if err != nil {
			return nil, err
		}
		if u == nil {
			return nil, nil
		}
		return &tg.InputUser{
			UserID:     u.UserId,
			AccessHash: u.AccessHash,
		}, nil
	}
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

func SaveUser(id int64, hash int64) error {
	var u *UserInfo
	userMapMutex.Lock()
	u = userDbMap[id]
	userMapMutex.Unlock()
	if u == nil || u.AccessHash != hash {
		if u == nil {
			u = &UserInfo{
				UserId:     id,
				AccessHash: hash,
			}
		}

		if wotoValues.IsRealOwner(id) {
			u.Permission = wotoGlobals.Owner
		}

		return NewUser(u)
	}
	return nil
}

func NewUser(u *UserInfo) error {
	dbMutex.Lock()
	tx := dbSession.Begin()
	tx.Save(u)
	tx.Commit()
	dbMutex.Unlock()
	u.SetCacheTime()
	userMapMutex.Lock()
	userDbMap[u.UserId] = u
	userMapMutex.Unlock()
	return nil
}

func GetUserFromId(id int64) (*UserInfo, error) {
	if dbSession == nil {
		return nil, ErrNoSession
	}

	userMapMutex.Lock()
	u := userDbMap[id]
	userMapMutex.Unlock()
	if u != nil {
		u.SetCacheTime()
		return u, nil
	}

	u = &UserInfo{}
	dbMutex.Lock()
	dbSession.Model(modelUser).Where("user_id = ?", id).Take(u)
	dbMutex.Unlock()
	if u.UserId != id || u.AccessHash == 0 {
		if wotoValues.IsRealOwner(id) {
			u.Permission = wotoGlobals.Owner
			u.AccessHash = wotoGlobals.Self.AccessHash
			u.UserId = id
			err := NewUser(u)
			if err != nil {
				return nil, err
			}
			return u, nil
		}
		// not found
		return nil, nil
	}

	u.SetCacheTime()
	userMapMutex.Lock()
	userDbMap[u.UserId] = u
	userMapMutex.Unlock()

	return u, nil
}
