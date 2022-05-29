package statsPlugin

import "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"

func LoadAllHandlers(manager *entryManager.EntryManager) {
	manager.AddHandlers(statusCmd, statusHandler)
}
