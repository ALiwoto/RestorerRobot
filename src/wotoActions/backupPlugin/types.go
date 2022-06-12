package backupPlugin

import (
	"sync"
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
)

type BackupScheduleContainer struct {
	DatabaseConfig *wotoConfig.ValueSection
	LastBackupDate time.Time
	BackupInterval time.Duration
	currentInfo    *wotoGlobals.BackupInfo
	ChatIDs        []int64
	mut            *sync.Mutex
	isSleeping     bool
}

type BackupScheduleManager struct {
	containers    []*BackupScheduleContainer
	checkInterval time.Duration
	ChatIDs       []int64
}
