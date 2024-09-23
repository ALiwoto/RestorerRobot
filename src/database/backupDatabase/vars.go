package backupDatabase

import (
	"sync"

	wg "github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/ALiwoto/ssg/ssg"
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
