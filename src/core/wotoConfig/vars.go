package wotoConfig

import "github.com/AnimeKaizoku/ssg/ssg/strongParser"

var WotoConf *strongParser.MainAndArrayContainer[MainConfigSection, ValueSection]

// _backupTypesMap is a read-only map that contains the correct types of
// backup, if this map returns false for a backup-type, it means that backup-type
// is either invalid or has incorrect spelling.
var _backupTypesMap = map[DatabaseBackupType]bool{
	BackupTypeDump:      true,
	BackupTypeSQL:       true,
	BackupTypeSQLite:    true,
	BackupTypeFile:      true,
	BackupTypeDirectory: true,
	BackupTypeFolder:    true,
}
