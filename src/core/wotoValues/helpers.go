package wotoValues

import (
	em "github.com/ALiwoto/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/gotd/td/tg"
)

func AppendEntryManager(manager ...*em.EntryManager) {
	EntryMaster = append(EntryMaster, manager...)
}

func IsRealOwner(id int64) bool {
	if wotoGlobals.Self == nil || wotoGlobals.Self.Bot {
		return false
	}

	return id == wotoGlobals.Self.ID
}

func SetSelf(s *tg.User) {
	wotoGlobals.Self = s
}
