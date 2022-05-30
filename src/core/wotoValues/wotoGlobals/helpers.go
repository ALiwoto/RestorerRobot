package wotoGlobals

import "strings"

func IsDatabaseUrl(value string) bool {
	return strings.HasPrefix(value, "postgresql://")
}
