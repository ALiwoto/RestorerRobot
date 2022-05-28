package wotoValues

import (
	em "github.com/ALiwoto/wotoub/wotoub/core/wotoEntry/enteryManager"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoValues/wotoGlobals"
	"github.com/gotd/td/tg"
)

func AppendEnteryManager(manager ...*em.EnteryManager) {
	EnetryMaster = append(EnetryMaster, manager...)
}

func IsRealOwner(id int64) bool {
	return wotoGlobals.Self != nil && id == wotoGlobals.Self.ID
}

func SetSelf(s *tg.User) {
	wotoGlobals.Self = s
}
