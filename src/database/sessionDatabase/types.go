package sessionDatabase

import (
	wg "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
)

type UserInfo struct {
	UserId     int64             `json:"user_id" gorm:"primaryKey"`
	AccessHash int64             `json:"access_hash"`
	Permission wg.UserPermission `json:"permission"`
}
