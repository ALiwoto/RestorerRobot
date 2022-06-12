package wotoGlobals

import (
	"strings"

	"github.com/AnimeKaizoku/ssg/ssg"
)

// --------------------------------------------------------

func (i *PeerInfo) GetRealID() int64 {
	if i.PeerType == PeerTypeChannel {
		idStr := ssg.ToBase10(i.PeerId)
		if strings.HasPrefix(idStr, "-100") {
			idStr = strings.TrimPrefix(idStr, "-100")
			return ssg.ToInt64(idStr)
		}
	}

	return i.PeerId
}

// --------------------------------------------------------

func (s BackupStatus) IsPending() bool {
	return s == BackupStatusPending
}

func (s BackupStatus) IsCanceled() bool {
	return s == BackupStatusCanceled
}

func (s BackupStatus) IsFinished() bool {
	return s == BackupStatusFinished
}

func (s BackupStatus) IsFailed() bool {
	return s == BackupStatusFailed
}

func (s BackupStatus) IsUnknown() bool {
	return s == BackupStatusUnknown
}

//---------------------------------------------------------

func (i *BackupInfo) GetStrStatus() string {
	switch i.Status {
	case BackupStatusPending:
		return "Pending"
	case BackupStatusCanceled:
		return "Canceled"
	case BackupStatusFinished:
		return "Finished"
	case BackupStatusFailed:
		return "Failed"
	case BackupStatusUnknown:
		fallthrough
	default:
		return "Unknown"
	}
}

func (i *BackupInfo) SetAsPending() {
	i.Status = BackupStatusPending
}

func (i *BackupInfo) SetAsCanceled() {
	i.Status = BackupStatusCanceled
}

func (i *BackupInfo) SetAsFinished() {
	i.Status = BackupStatusFinished
}

func (i *BackupInfo) SetAsFailed() {
	i.Status = BackupStatusFailed
}
