package database

import "github.com/AnimeKaizoku/RestorerRobot/src/database/sessionDatabase"

func StartDatabase() error {
	err := sessionDatabase.StartDatabase()
	if err != nil {
		return err
	}

	return nil
}
