package backupDatabase

import (
	"sync"

	wg "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/AnimeKaizoku/ssg/ssg"
	"gorm.io/gorm"
)

var (
	ModelDatabaseInfo = &wg.DataBaseInfo{}
	ModelBackupInfo   = &wg.BackupInfo{}
)

var (
	dbSession *gorm.DB
	dbMutex   *sync.Mutex
)

var (
	databaseInfoMap = ssg.NewSafeMap[string, wg.DataBaseInfo]()
	backupInfoMap   = ssg.NewSafeMap[wg.BackupUniqueIdValue, wg.BackupInfo]()
)
