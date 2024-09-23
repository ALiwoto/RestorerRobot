package helpPlugin

import "github.com/ALiwoto/RestorerRobot/src/core/wotoEntry/entryManager"

func LoadAllHandlers(manager *entryManager.EntryManager) {
	manager.AddHandlers(configsCmd, configsHandler)
	manager.AddHandlers(startCmd, startHandler)
	manager.AddHandlers(dbInfoCmd, dbInfoHandler)
}
