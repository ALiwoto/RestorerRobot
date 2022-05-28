package wotoActions

import (
	"github.com/ALiwoto/wotoub/wotoub/core/wotoConfig"
	em "github.com/ALiwoto/wotoub/wotoub/core/wotoEntry/enteryManager"
	wv "github.com/ALiwoto/wotoub/wotoub/core/wotoValues"
	"github.com/ALiwoto/wotoub/wotoub/wotoActions/funPlugin"
	"github.com/ALiwoto/wotoub/wotoub/wotoActions/statsPlugin"
)

func LoadAllHandlers() {
	loadManagers()

	statsPlugin.LoadAllHandlers(wv.CommandManager)
	funPlugin.LoadAllHandlers(wv.CommandManager)
}

func loadManagers() {
	wv.CommandManager = em.NewManager(wotoConfig.WotoConf.CmdPrefixes)
	wv.AppendEnteryManager(wv.CommandManager)
}
