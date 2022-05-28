package sessionDatabase

import (
	"time"

	wg "github.com/ALiwoto/wotoub/wotoub/core/wotoValues/wotoGlobals"
)

type UserInfo struct {
	UserId      int64             `json:"user_id" gorm:"primaryKey"`
	AccessHash  int64             `json:"access_hash"`
	Permission  wg.UserPermission `json:"permission"`
	ShouldStalk bool              `json:"should_stalk"`
	cachedTime  time.Time         `json:"-" gorm:"-" sql:"-"`
}
