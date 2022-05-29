package funPlugin

import "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"

func LoadAllHandlers(manager *entryManager.EntryManager) {
	manager.AddHandlers(CmdMention, mentionHandler)
	manager.AddHandlers(CmdFish, fishHandler)
}
