package enteryManager

import (
	"sync"

	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/tg"
)

type MessageHandler func(container *WotoContainer) error
type MessageFilter func(message *tg.Message) bool
type AllowableUserFilter func(id int64) bool

type ManagerGroup []*EnteryManager

type EnteryManager struct {
	triggers    []rune
	enteryMap   map[string]*Entery
	enteryMutex *sync.Mutex
}

type Entery struct {
	enabled       bool
	restrictUsers bool
	UniqueName    string
	handlers      []MessageHandler
	allowedUsers  AllowableUserFilter
	// internalCondition is only for internal usage. determining commands.
	internalCondition MessageFilter
	// filters field. all filters should return if handlers are going to run.
	filters []MessageFilter
}

type WotoContainer struct {
	//#region: Message fields

	OriginMessage  tg.MessageClass
	ServiceMessage *tg.MessageService
	Message        *tg.Message

	//#endregion: Message fields

	//#region: Update fields
	Update                    tg.UpdateClass
	UpdateNewScheduledMessage *tg.UpdateNewScheduledMessage
	UpdateNewChannelMessage   *tg.UpdateNewChannelMessage
	UpdateNewMessage          *tg.UpdateNewMessage
	UpdateEditMessage         *tg.UpdateEditMessage

	//#endregion: Update fields

	Entities   *tg.Entities
	Answerable message.AnswerableMessageUpdate

	//#region: private fields

	//#endregion: private fields
}
