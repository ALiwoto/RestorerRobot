package backupPlugin

import "time"

const (
	forceBackupCmd = "forceBackup"
)

const (
	managerTimeInterval = 4 * time.Hour
)

const (
	MinSchedulerDuration = managerTimeInterval
)
