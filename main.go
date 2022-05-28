package main

import (
	"github.com/ALiwoto/wotoub/wotoub/core/utils/logging"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoAuth"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoConfig"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoEntry"
	"github.com/ALiwoto/wotoub/wotoub/database"
)

func main() {
	f := logging.LoadLogger(true)
	if f != nil {
		defer f()
	}

	runApp()
}

func runApp() {
	err := wotoConfig.LoadConfig()
	if err != nil {
		logging.Fatal("Error loading config:", err.Error())
	}

	err = database.StartDatabase()
	if err != nil {
		logging.Fatal("Error starting database:", err.Error())
	}

	u := wotoEntry.MainUpdateEntry
	err = wotoAuth.AuthorizeClient(wotoConfig.WotoConf.Phone, u)
	if err != nil {
		logging.Fatal("Error authorizing client:", err.Error())
	}
}
