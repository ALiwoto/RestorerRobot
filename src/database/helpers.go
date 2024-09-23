package database

import (
	"github.com/ALiwoto/RestorerRobot/src/core/utils/logging"
	"github.com/ALiwoto/RestorerRobot/src/core/wotoConfig"
	"github.com/ALiwoto/RestorerRobot/src/database/backupDatabase"
	"github.com/ALiwoto/RestorerRobot/src/database/sessionDatabase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

	db, err = gorm.Open(sqlite.Open(wotoConfig.GetDbPath()), conf)
	if err != nil {
		return err
	}

	dbSession = db

	logging.Info("Database connected")

	// Create tables if they don't exist
	err = dbSession.AutoMigrate(
		sessionDatabase.ModelUser,
		backupDatabase.ModelBackupInfo,
		backupDatabase.ModelDatabaseInfo,
	)
	if err != nil {
		return err
	}

	err = sessionDatabase.StartDatabase(dbSession, dbMutex)
	if err != nil {
		return err
	}

	err = backupDatabase.StartDatabase(dbSession, dbMutex)
	if err != nil {
		return err
	}

	return nil
}
