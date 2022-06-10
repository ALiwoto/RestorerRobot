package backupDatabase

import (
	"sync"
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	wg "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"gorm.io/gorm"
)

func StartDatabase(db *gorm.DB, mut *sync.Mutex) error {
	dbSession = db
	dbMutex = mut

	err := forceImportSections()
	if err != nil {
		return err
	}

	return nil
}

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

func GetLastBackupDate(configName string) time.Time {
	info := GetDatabaseInfo(configName)
	if info == nil {
		return time.Time{}
	}

	return info.LastBackup
}

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

func GetBackupInfo(uniqueId wg.BackupUniqueIdValue) *wg.BackupInfo {
	info := backupInfoMap.Get(uniqueId)
	if info != nil {
		return info
	}

	info = &wg.BackupInfo{}
	dbMutex.Lock()
	dbSession.Model(ModelDatabaseInfo).Where("backup_uniqueId = ?", uniqueId).Take(info)
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

func GenerateBackupInfo(fromNow time.Duration, by int64) *wg.BackupInfo {
	info := &wg.BackupInfo{
		BackupUniqueId: wg.GenerateBackupUniqueId(),
		BackupDate:     time.Now().Add(fromNow),
		RequestedBy:    by,
	}

	NewBackupInfo(info)

	return info

}
