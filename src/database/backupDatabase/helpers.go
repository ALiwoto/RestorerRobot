package backupDatabase

import (
	"sync"
	"time"

	"github.com/ALiwoto/RestorerRobot/src/core/wotoConfig"
	"github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
	wg "github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"gorm.io/gorm"
)

// StartDatabase will initialize the variables for backupDatabase package.
func StartDatabase(db *gorm.DB, mut *sync.Mutex) error {
	dbSession = db
	dbMutex = mut

	err := forceImportSections()
	if err != nil {
		return err
	}

	return nil
}

// forceImportSections will forcefully import all sections into the database.
func forceImportSections() error {
	sections := wotoConfig.WotoConf.Sections
	if len(sections) == 0 {
		return nil
	}

	var err error

	for _, currentSection := range sections {
		currentInfo := GetDatabaseInfo(currentSection.GetSectionName())
		if currentInfo == nil {
			err = NewDatabaseInfo(&wg.DataBaseInfo{
				DatabaseName: currentSection.GetSectionName(),
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetLastBackupDate will return the last time the specified database was backed up.
func GetLastBackupDate(configName string) time.Time {
	info := GetDatabaseInfo(configName)
	if info == nil {
		return time.Time{}
	}

	return info.LastBackup
}

// GetDatabaseInfo returns the database info using its specified name.
func GetDatabaseInfo(name string) *wg.DataBaseInfo {
	info := databaseInfoMap.Get(name)
	if info != nil {
		return info
	}

	info = &wg.DataBaseInfo{}
	dbMutex.Lock()
	dbSession.Model(ModelDatabaseInfo).Where("database_name = ?", name).Take(info)
	dbMutex.Unlock()

	if info.DatabaseName != name {
		// not found
		return nil
	}

	databaseInfoMap.Add(name, info)

	return info
}

func GetLastBackupStatus(name string) wg.BackupStatus {
	dbInfo := GetDatabaseInfo(name)
	if dbInfo == nil || dbInfo.LastBackupUniqueId.IsInvalid() {
		return wg.BackupStatusUnknown
	}

	backupInfo := GetBackupInfo(dbInfo.LastBackupUniqueId)
	if backupInfo == nil {
		return wg.BackupStatusUnknown
	}

	return backupInfo.Status
}

// GetBackupFinishedCount returns the count of the finished backups for the
// specified database name.
func GetBackupFinishedCount(name string) int64 {
	var count int64
	dbMutex.Lock()
	m := dbSession.Model(ModelBackupInfo)
	m.Where("database_name = ? AND status = ?", name, wotoGlobals.BackupStatusFinished).Count(&count)
	dbMutex.Unlock()
	return count
}

func NewDatabaseInfo(info *wg.DataBaseInfo) error {
	dbMutex.Lock()
	tx := dbSession.Begin()
	tx.Save(info)
	tx.Commit()
	dbMutex.Unlock()
	databaseInfoMap.Add(info.DatabaseName, info)
	return nil
}

func UpdateDatabaseInfo(info *wg.DataBaseInfo) error {
	dbMutex.Lock()
	tx := dbSession.Begin()
	tx.Save(info)
	tx.Commit()
	dbMutex.Unlock()
	return nil
}

func UpdateBackupInfo(info *wg.BackupInfo) error {
	dbMutex.Lock()
	tx := dbSession.Begin()
	tx.Save(info)
	tx.Commit()
	dbMutex.Unlock()
	return nil
}

func GetBackupInfo(uniqueId wg.BackupUniqueIdValue) *wg.BackupInfo {
	info := backupInfoMap.Get(uniqueId)
	if info != nil {
		return info
	}

	info = &wg.BackupInfo{}
	dbMutex.Lock()
	dbSession.Model(ModelBackupInfo).Where("backup_unique_id = ?", uniqueId).Take(info)
	dbMutex.Unlock()

	if info.BackupUniqueId != uniqueId {
		// not found
		return nil
	}

	backupInfoMap.Add(uniqueId, info)

	return info
}

func NewBackupInfo(info *wg.BackupInfo) error {
	dbMutex.Lock()
	tx := dbSession.Begin()
	tx.Save(info)
	tx.Commit()
	dbMutex.Unlock()
	backupInfoMap.Add(info.BackupUniqueId, info)
	return nil
}

func GenerateBackupInfo(name string, fromNow time.Duration, by int64) *wg.BackupInfo {
	info := &wg.BackupInfo{
		BackupUniqueId: wg.GenerateBackupUniqueId(),
		DatabaseName:   name,
		BackupDate:     time.Now().Add(fromNow),
		RequestedBy:    by,
		Status:         wg.BackupStatusPending,
	}

	NewBackupInfo(info)

	return info

}
