package wotoGlobals

import "time"

type UserPermission int
type PeerType int
type BackupUniqueIdValue string

type PeerInfo struct {
	PeerId     int64    `json:"peer_id" gorm:"primaryKey"`
	AccessHash int64    `json:"access_hash"`
	PeerType   PeerType `json:"peer_type"`
}

// DataBaseInfo contains information about a database with its
// name using as the primary key.
type DataBaseInfo struct {
	DatabaseName       string `gorm:"primaryKey"`
	LastBackup         time.Time
	LastBackupUniqueId BackupUniqueIdValue
}

// BackupInfo contains the information about each backup the bot gets.
type BackupInfo struct {
	BackupUniqueId BackupUniqueIdValue `gorm:"primaryKey"`
	BackupDate     time.Time
	RequestedBy    int64
}
