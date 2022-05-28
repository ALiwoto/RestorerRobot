package database

import "github.com/ALiwoto/wotoub/wotoub/database/sessionDatabase"

func StartDatabase() error {
	err := sessionDatabase.StartDatabase()
	if err != nil {
		return err
	}

	return nil
}
