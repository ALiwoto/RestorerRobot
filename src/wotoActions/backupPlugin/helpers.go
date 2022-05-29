package backupPlugin

import "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"

func LoadAllHandlers(manager *entryManager.EntryManager) {
	manager.AddHandlers(forceBackupCmd, forceBackupHandler)

	loadScheduler()
}

func loadScheduler() {
	// TODO: implement
}
