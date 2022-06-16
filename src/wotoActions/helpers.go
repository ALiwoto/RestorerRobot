package wotoActions

import (
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	em "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	wv "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues"
	"github.com/AnimeKaizoku/RestorerRobot/src/wotoActions/backupPlugin"
	"github.com/AnimeKaizoku/RestorerRobot/src/wotoActions/helpPlugin"
	"github.com/AnimeKaizoku/RestorerRobot/src/wotoActions/statsPlugin"
)

// LoadAllHandlers will load all handlers from all plugins.
// WARNING: helpPlugin imports backupPlugin, do NOT import any other
// plugin inside of backupPlugin to prevent cycle import error.
// (importing plugins inside of plugins is a big mistake by itself,
// but this time I was just too lazy to move types and methods to
// wotoValues).
func LoadAllHandlers() {
	loadManagers()

	statsPlugin.LoadAllHandlers(wv.CommandManager)
	backupPlugin.LoadAllHandlers(wv.CommandManager)
	helpPlugin.LoadAllHandlers(wv.CommandManager)
}

func loadManagers() {
	wv.CommandManager = em.NewManager(wotoConfig.GetPrefixes())
	wv.AppendEntryManager(wv.CommandManager)
}
