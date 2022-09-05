package wotoGlobals

import (
	"strings"
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
)

func IsPostgresDatabaseUrl(value string) bool {
	return strings.HasPrefix(value, "postgresql://") || strings.HasPrefix(value, "postgres")
}

func GenerateBackupUniqueId() BackupUniqueIdValue {
	strValue := backupUniqueIdPrefix
	strValue += ssg.ToBase32(time.Now().Unix()) + ssg.ToBase10(backupUniqueIdNumGenerator.Next())

	return BackupUniqueIdValue(strValue)
}

func IsBackupUniqueId(value string) bool {
	return strings.HasPrefix(value, backupUniqueIdPrefix)
}
