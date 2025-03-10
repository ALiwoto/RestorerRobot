package backupPlugin

import (
	"sync"

	"github.com/ALiwoto/RestorerRobot/src/core/wotoConfig"
	"github.com/ALiwoto/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/ALiwoto/RestorerRobot/src/database/backupDatabase"
)

func GetContainerByName(name string) *BackupScheduleContainer {
	if scheduleManager == nil {
		return nil
	}

	return scheduleManager.GetContainerByName(name)
}

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
		containers:    make([]*BackupScheduleContainer, len(configs)),
		checkInterval: wotoConfig.GetScheduleManagerInterval(),
		ChatIDs:       wotoConfig.GetGlobalLogChannels(),
	}

	containersMutex := &sync.Mutex{}

	for i := range configs {
		manager.containers[i] = &BackupScheduleContainer{
			DatabaseConfig: configs[i],
			LastBackupDate: backupDatabase.GetLastBackupDate(configs[i].GetSectionName()),
			BackupInterval: manager.convertToBackupInterval(configs[i].BackupInterval),
			ChatIDs:        manager.ChatIDs,
			mut:            containersMutex,
		}
	}

	scheduleManager = manager
	go manager.RunChecking()
}

func toBackupType(value string) wotoConfig.DatabaseBackupType {
	return wotoConfig.DatabaseBackupType(value)
}
