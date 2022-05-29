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

//  case *tg.InputPeerEmpty: // inputPeerEmpty#7f3b18ea
//  case *tg.InputPeerSelf: // inputPeerSelf#7da07ec9
//  case *tg.InputPeerChat: // inputPeerChat#35a95cb9
//  case *tg.InputPeerUser: // inputPeerUser#dde8a54c
//  case *tg.InputPeerChannel: // inputPeerChannel#27bcbbfc
//  case *tg.InputPeerUserFromMessage: // inputPeerUserFromMessage#a87b0a1c
//  case *tg.InputPeerChannelFromMessage: // inputPeerChannelFromMessage#bd2a0840
const (
	PeerTypeEmpty = iota
	PeerTypeSelf
	PeerTypeChat
	PeerTypeUser
	PeerTypeChannel
	PeerTypeUserFromMessage
	PeerTypeChannelFromMessage
)
