package backupUtils

import "time"

type GenerateCaptionOptions struct {
	ConfigName     string
	BackupInitType string
	InitiatedBy    string
	UserId         int64
	DateTime       time.Time
	FileSize       string
	BackupFormat   string
}
