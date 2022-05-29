package wotoActions

import (
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	em "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	wv "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues"
	"github.com/AnimeKaizoku/RestorerRobot/src/wotoActions/statsPlugin"
)

func LoadAllHandlers() {
	loadManagers()

	statsPlugin.LoadAllHandlers(wv.CommandManager)
}

func loadManagers() {
	wv.CommandManager = em.NewManager(wotoConfig.GetPrefixes())
	wv.AppendEntryManager(wv.CommandManager)
}
