package sessionDatabase

import (
	"errors"
	"sync"

	"github.com/AnimeKaizoku/ssg/ssg"
	"gorm.io/gorm"
)

var (
	dbSession *gorm.DB
	modelUser = &UserInfo{}
	dbMutex   = &sync.Mutex{}
	userDbMap = ssg.NewSafeMap[int64, UserInfo]()
)

var (
	ErrNoSession      = errors.New("database session is not initialized")
	ErrTooManyRevokes = errors.New("token has been revoked too many times")
)
