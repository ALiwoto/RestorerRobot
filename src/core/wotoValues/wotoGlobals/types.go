package wotoGlobals

type UserPermission int
type PeerType int

type PeerInfo struct {
	PeerId     int64    `json:"peer_id" gorm:"primaryKey"`
	AccessHash int64    `json:"access_hash"`
	PeerType   PeerType `json:"peer_type"`
}
