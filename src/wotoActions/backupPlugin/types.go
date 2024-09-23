package backupPlugin

import (
	"sync"
	"time"

	"github.com/ALiwoto/RestorerRobot/src/core/wotoConfig"
	"github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
)

type BackupScheduleContainer struct {
	// DatabaseConfig is the config section of the target db belonging to
	// this container.
	DatabaseConfig *wotoConfig.ValueSection

	// LastBackupDate is the date of the last time we got backup from this db.
	LastBackupDate time.Time

	// BackupInterval is the interval that we will get backup.
	BackupInterval time.Duration

	// currentInfo fields contains the current BackupInfo of this container.
	currentInfo *wotoGlobals.BackupInfo

	// ChatIDs field specifies the global log channels that we have
	// to send the backup to.
	ChatIDs []int64

	// mut is a shared mutex between all related BackupScheduleContainer values
	// (normally, all values, unless mentioned otherwise.)
	mut *sync.Mutex

	// isSleeping fields specifies if the current container is in sleeping phase.
	isSleeping bool
}

type BackupScheduleManager struct {
	// containers is an array of all existing BackupScheduleContainer that
	// needs to be checked and handled by this BackupScheduleManager value.
	containers []*BackupScheduleContainer

	// checkInterval is the interval of each for loop of checker.
	checkInterval time.Duration

	// ChatIDs field specifies the global log channels that we have
	// to send the backup to.
	ChatIDs []int64
}
