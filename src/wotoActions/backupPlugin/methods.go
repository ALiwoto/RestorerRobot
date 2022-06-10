package backupPlugin

import "time"

// --------------------------------------------------------

func (m *BackupScheduleManager) Run() {

}

func (m *BackupScheduleManager) convertToBackupInterval(days int) time.Duration {
	return time.Duration(days) * 24 * time.Hour
}

// --------------------------------------------------------

func (m *BackupScheduleContainer) IsWithin(d time.Duration) bool {
	return time.Since(m.LastBackupDate)-m.BackupInterval >= d
}

// --------------------------------------------------------
// --------------------------------------------------------
