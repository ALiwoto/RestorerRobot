package sessionDatabase

import (
	"errors"
	"sync"

	"github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/ALiwoto/ssg/ssg"
	"gorm.io/gorm"
)

var (
	dbSession *gorm.DB
	ModelUser = &wotoGlobals.PeerInfo{}
	dbMutex   = &sync.Mutex{}
	peerDbMap = ssg.NewSafeMap[int64, wotoGlobals.PeerInfo]()
)

var (
	ErrNoSession      = errors.New("database session is not initialized")
	ErrTooManyRevokes = errors.New("token has been revoked too many times")
)
