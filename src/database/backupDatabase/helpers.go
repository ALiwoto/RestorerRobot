package backupDatabase

import (
	"time"

	"gorm.io/gorm"
)

func StartDatabase(db *gorm.DB) error {
	dbSession = db
	return nil
}

func GetLastBackupDate(configName string) time.Time {

	return time.Time{}
}
