package backupPlugin

import (
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/AnimeKaizoku/RestorerRobot/src/database/backupDatabase"
)

func LoadAllHandlers(manager *entryManager.EntryManager) {
	manager.AddHandlers(forceBackupCmd, forceBackupHandler)

	loadScheduler()
}

func loadScheduler() {
	if scheduleManager != nil {
		return
	}

	configs := wotoConfig.GetDatabasesConfigs()
	if len(configs) == 0 {
		scheduleManager = &BackupScheduleManager{
			containers: nil,
		}
		return
	}

	manager := &BackupScheduleManager{
		containers: make([]BackupScheduleContainer, len(configs)),
	}

	for i := 0; i < len(configs); i++ {
		manager.containers[i] = BackupScheduleContainer{
			DatabaseConfig: configs[i],
			LastBackupDate: backupDatabase.GetLastBackupDate(configs[i].GetSectionName()),
			BackupInterval: manager.convertToBackupInterval(configs[i].BackupInterval),
		}
	}

	scheduleManager = manager
	go manager.Run()
}
