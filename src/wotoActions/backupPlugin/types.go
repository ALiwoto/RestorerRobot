package backupPlugin

import (
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
)

type BackupScheduleContainer struct {
	DatabaseConfig *wotoConfig.ValueSection
	LastBackupDate time.Time
	BackupInterval time.Duration
}

type BackupScheduleManager struct {
	containers []BackupScheduleContainer
}
