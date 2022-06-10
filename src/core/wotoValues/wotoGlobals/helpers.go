package wotoGlobals

import (
	"strings"
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
)

func IsDatabaseUrl(value string) bool {
	return strings.HasPrefix(value, "postgresql://")
}

func GenerateBackupUniqueId() BackupUniqueIdValue {
	strValue := backupUniqueIdPrefix
	strValue += ssg.ToBase32(time.Now().Unix()) + ssg.ToBase10(backupUniqueIdNumGenerator.Next())

	return BackupUniqueIdValue(strValue)
}
