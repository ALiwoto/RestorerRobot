package statsPlugin

import "github.com/ALiwoto/wotoub/wotoub/core/wotoEntry/enteryManager"

func LoadAllHandlers(manager *enteryManager.EnteryManager) {
	manager.AddHandlers(CmdWotoUb, wotoUbHandler)
}
