package backupPlugin

import "errors"

var scheduleManager *BackupScheduleManager

var (
	ErrNoPathOrUrlSet = errors.New("the section doesn't have any path or url set")
)
