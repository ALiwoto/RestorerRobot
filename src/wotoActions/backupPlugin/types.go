package backupPlugin

import (
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
)

type BackupScheduleContainer struct {
	DatabaseConfig *wotoConfig.ValueSection
	LastBackupDate time.Time
}

type BackupScheduleManager struct {
	containers []BackupScheduleContainer
}
