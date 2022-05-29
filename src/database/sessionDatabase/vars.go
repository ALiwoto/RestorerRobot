package sessionDatabase

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

var (
	dbSession    *gorm.DB
	modelUser    = &UserInfo{}
	dbMutex      = &sync.Mutex{}
	userMapMutex = &sync.Mutex{}
	userDbMap    = make(map[int64]*UserInfo)
)

var (
	ErrNoSession      = errors.New("database session is not initialized")
	ErrTooManyRevokes = errors.New("token has been revoked too many times")
)
