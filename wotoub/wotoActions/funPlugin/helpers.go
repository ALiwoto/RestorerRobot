package funPlugin

import "github.com/ALiwoto/wotoub/wotoub/core/wotoEntry/enteryManager"

func LoadAllHandlers(manager *enteryManager.EnteryManager) {
	manager.AddHandlers(CmdMention, mentionHandler)
	manager.AddHandlers(CmdFish, fishHandler)
}
