package wotoGlobals

const (
	AppVersion = "1.0.0"
)

const (
	NormalUser UserPermission = iota
	Stalk
	Special
	Friend
	PseudoSudo
	Sudo
	Owner
)

const (
	backupUniqueIdPrefix = "backupZZ"
)

const (
	// BackupStatusUnknown is only there as a default and invalid
	// backup status, the Status field of `BackupInfo` struct should not
	// contain this value at all, if it does, user will just face an
	// unknown situation (which is considered as a bug and has to be fixed).
	BackupStatusUnknown BackupStatus = iota
	// BackupStatusPending, as it sounds like, is when the backup is
	// pending (it's within the range of our BackupScheduleManager
	// time interval and it's sleeping till it gets the right time to
	// trigger its backup), or maybe it's in the middle of backup process,
	// or zipping the file, or uploading the files, etc etc...
	BackupStatusPending
	// BackupStatusCanceled is when the backup is canceled by user.
	// Please notice that only this time will get canceled by user,
	// the next time we reach the right time, we have to start backup
	//  process. (if users want to completely remove the backup process
	//  of a database, they have to remove the section from the config file).
	BackupStatusCanceled
	// BackupStatusFailed is when there has been an error in the backup
	// process, those errors have to be reported to users and inserted
	//  into database as past records, so when users want to see the
	// backup status in the future, they get access to it.
	BackupStatusFailed
	// BackupStatusFinished is when the backup process has been
	// finished successfully without encountering any errors in the
	// middle of way.
	BackupStatusFinished
)

// case *tg.InputPeerEmpty: // inputPeerEmpty#7f3b18ea
// case *tg.InputPeerSelf: // inputPeerSelf#7da07ec9
// case *tg.InputPeerChat: // inputPeerChat#35a95cb9
// case *tg.InputPeerUser: // inputPeerUser#dde8a54c
// case *tg.InputPeerChannel: // inputPeerChannel#27bcbbfc
// case *tg.InputPeerUserFromMessage: // inputPeerUserFromMessage#a87b0a1c
// case *tg.InputPeerChannelFromMessage: // inputPeerChannelFromMessage#bd2a0840
const (
	PeerTypeEmpty = iota
	PeerTypeSelf
	PeerTypeChat
	PeerTypeUser
	PeerTypeChannel
	PeerTypeUserFromMessage
	PeerTypeChannelFromMessage
)

const (
	BaseUploadGoroutines = 20
)
