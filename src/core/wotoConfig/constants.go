package wotoConfig

// backup types
const (
	// BackupTypeDump will make a .dump file output from the postgres database.
	BackupTypeDump DatabaseBackupType = "dump"
	// BackupTypeSQL will make a .sql file output from the postgres database.
	BackupTypeSQL DatabaseBackupType = "sql"
	// BackupTypeSQLite will just make a zip from the sqlite file.
	BackupTypeSQLite DatabaseBackupType = "sqlite"
	// BackupTypeFile will just convert the file to zip.
	BackupTypeFile DatabaseBackupType = "file"
	// BackupTypeDirectory will just convert the directory to zip.
	BackupTypeDirectory DatabaseBackupType = "directory"
	// BackupTypeFolder will just convert the folder to zip.
	BackupTypeFolder DatabaseBackupType = "folder" // same as directory
)

const (
	CompressedFileExtension = ".zip"
)
