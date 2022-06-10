package backupDatabase

import (
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"gorm.io/gorm"
)

var (
	ModelDatabaseInfo = &wotoGlobals.DataBaseInfo{}
	ModelBackupInfo   = &wotoGlobals.BackupInfo{}

	dbSession *gorm.DB
)
