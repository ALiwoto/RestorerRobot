package wotoGlobals

import (
	"os"
	"strings"
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
)

func IsPostgresDatabaseUrl(value string) bool {
	return strings.HasPrefix(value, "postgresql://") || strings.HasPrefix(value, "postgres")
}

// IsValidLocalFileOrDir function will return true if and only if the value passed
// as argument to it represents an existing (and correct), local file/directory path.
func IsValidLocalFileOrDir(value string) bool {
	var err error
	if _, err = os.ReadDir(value); err == nil {
		return true
	}

	if _, err = os.ReadFile(value); err == nil {
		return true
	}

	return false
}

func GenerateBackupUniqueId() BackupUniqueIdValue {
	strValue := backupUniqueIdPrefix
	strValue += ssg.ToBase32(time.Now().Unix()) + ssg.ToBase10(backupUniqueIdNumGenerator.Next())

	return BackupUniqueIdValue(strValue)
}

func IsBackupUniqueId(value string) bool {
	return strings.HasPrefix(value, backupUniqueIdPrefix)
}
