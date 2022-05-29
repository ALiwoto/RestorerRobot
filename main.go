package main

import (
	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/logging"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoAuth"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry"
	"github.com/AnimeKaizoku/RestorerRobot/src/database"
	"github.com/go-faster/errors"
)

func main() {
	f := logging.LoadLogger(true)
	if f != nil {
		defer f()
	}

	err := runApp()
	if err != nil {
		logging.Fatal("Error running app: ", err.Error())
	}
}

func runApp() error {
	err := wotoConfig.LoadConfig()
	if err != nil {
		return errors.Wrap(err, "LoadConfig")
	}

	err = database.StartDatabase()
	if err != nil {
		return errors.Wrap(err, "StartDatabase")
	}

	u := wotoEntry.MainUpdateEntry
	err = wotoAuth.AuthorizeClient(u)
	if err != nil {
		return errors.Wrap(err, "AuthorizeClient")
	}

	return nil
}
