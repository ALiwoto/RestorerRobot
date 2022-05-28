package enteryManager

import (
	"context"
	"strings"
	"sync"

	"github.com/ALiwoto/argparser/argparser"
	"github.com/ALiwoto/wotoub/wotoub/core/utils/tgUtils"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoConfig"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoStyle"
	wg "github.com/ALiwoto/wotoub/wotoub/core/wotoValues/wotoGlobals"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/tg"
)

//---------------------------------------------------------

func (m *EnteryManager) IsEmpty() bool {
	return len(m.enteryMap) == 0
}

func (m *EnteryManager) GetEntery(name string) *Entery {
	if m.enteryMutex == nil {
		m.enteryMutex = &sync.Mutex{}
	}
	m.enteryMutex.Lock()
	e := m.enteryMap[name]
	m.enteryMutex.Unlock()
	return e
}

func (m *EnteryManager) AddEntery(name string, e *Entery) {
	if m.enteryMutex == nil {
		m.enteryMutex = &sync.Mutex{}
	}
	m.enteryMutex.Lock()
	m.enteryMap[name] = e
	m.enteryMutex.Unlock()
}

func (m *EnteryManager) AddHandlers(name string, h ...MessageHandler) *Entery {
	e := &Entery{
		enabled:    true,
		handlers:   h,
		UniqueName: name,
	}

	e.internalCondition = func(message *tg.Message) bool {
		return len(message.Message) > 1 && strings.HasPrefix(message.Message[1:], name)
	}

	m.AddEntery(name, e)
	return e
}

func (m *EnteryManager) RemoveEntery(name string) {
	if m.enteryMutex == nil {
		m.enteryMutex = &sync.Mutex{}
	}
	m.enteryMutex.Lock()
	delete(m.enteryMap, name)
	m.enteryMutex.Unlock()
}

func (m *EnteryManager) AddTriggers(t ...rune) []rune {
	m.triggers = append(m.triggers, t...)
	return m.triggers
}

func (m *EnteryManager) GetTriggers() []rune {
	return m.triggers
}

func (m *EnteryManager) SetTriggers(t []rune) {
	m.triggers = t
}

func (m *EnteryManager) ShouldRevoke(message string) bool {
	if len(message) < 2 {
		return false
	}

	rMessage := []rune(message)
	for _, trigger := range m.triggers {
		if rMessage[0] == trigger {
			return true
		}
	}
	return false
}

func (m *EnteryManager) Revoke(container *WotoContainer) (next bool) {
	message := container.Message
	if m.ShouldRevoke(message.Message) {
		m.enteryMutex.Lock()
		for _, entery := range m.enteryMap {
			if entery.IsEnabled() && entery.ShouldRun(message) {
				next = entery.RunHandlers(container)
				if !next {
					break
				}
			}
		}
		m.enteryMutex.Unlock()
		return false
	}

	return true
}

//---------------------------------------------------------

func (e *Entery) IsEnabled() bool {
	return e.enabled
}

func (e *Entery) Disable() {
	e.enabled = false
}

func (e *Entery) Enable() {
	e.enabled = true
}

func (e *Entery) IsEmpty() bool {
	return len(e.handlers) == 0
}

func (e *Entery) AddHandlers(h ...MessageHandler) {
	e.handlers = append(e.handlers, h...)
}

func (e *Entery) GetHandlers() []MessageHandler {
	return e.handlers
}

func (e *Entery) SetHandlers(h []MessageHandler) {
	e.handlers = h
}

func (e *Entery) ShouldRun(message *tg.Message) bool {
	if e.internalCondition != nil {
		if !e.internalCondition(message) {
			return false
		}
	}

	if e.restrictUsers && e.allowedUsers != nil {
		if !e.allowedUsers(tgUtils.GetIdFromPeerClass(message.FromID)) {
			return false
		}
	}

	if len(e.filters) == 0 {
		return true
	}

	for _, current := range e.filters {
		if !current(message) {
			return false
		}
	}

	return true
}

func (e *Entery) RunHandlers(container *WotoContainer) (next bool) {
	var err error
	if len(e.handlers) == 0 {
		return true
	}
	for _, handler := range e.handlers {
		err = handler(container)
		if err != nil {
			if err == ErrEndGroups {
				return false
			}

			if err == ErrContinueGroups {
				return true
			}
		}
	}
	return true
}

//---------------------------------------------------------

func (g ManagerGroup) TryToRun(container *WotoContainer) {
	for _, manager := range g {
		if !manager.Revoke(container) {
			return
		}
	}
}

//---------------------------------------------------------

func (c *WotoContainer) GetAnswerable() message.AnswerableMessageUpdate {
	return c.Answerable
}

func (c *WotoContainer) Ctx() context.Context {
	return gCtx
}

func (c *WotoContainer) ReplyText(text string) (tg.UpdatesClass, error) {
	return c.GetReplyBuilder().Text(gCtx, text)
}

func (c *WotoContainer) ReplyStrikeText(text string) (tg.UpdatesClass, error) {
	return c.GetReplyBuilder().StyledText(gCtx, styling.Strike(text))
}

func (c *WotoContainer) ReplyBoldText(text string) (tg.UpdatesClass, error) {
	return c.GetReplyBuilder().StyledText(gCtx, styling.Bold(text))
}

func (c *WotoContainer) ReplyItalicText(text string) (tg.UpdatesClass, error) {
	return c.GetReplyBuilder().StyledText(gCtx, styling.Italic(text))
}

func (c *WotoContainer) ReplyMonoText(text string) (tg.UpdatesClass, error) {
	return c.GetReplyBuilder().StyledText(gCtx, styling.Code(text))
}

func (c *WotoContainer) ReplyCodeText(text string) (tg.UpdatesClass, error) {
	return c.GetReplyBuilder().StyledText(gCtx, styling.Code(text))
}

func (c *WotoContainer) ReplyStyledText(s wotoStyle.WStyle) (tg.UpdatesClass, error) {
	return c.GetReplyBuilder().StyledText(gCtx, s.GetStylingArray()...)
}

func (c *WotoContainer) ReplyTextf(format string, args ...interface{}) (tg.UpdatesClass, error) {
	return c.GetReplyBuilder().Textf(gCtx, format, args...)
}

func (c *WotoContainer) SendText(text string) (tg.UpdatesClass, error) {
	return c.GetSendBuilder().Text(gCtx, text)
}

func (c *WotoContainer) SendStrikeText(text string) (tg.UpdatesClass, error) {
	return c.GetSendBuilder().StyledText(gCtx, styling.Strike(text))
}

func (c *WotoContainer) SendSpoilerText(text string) (tg.UpdatesClass, error) {
	return c.GetSendBuilder().StyledText(gCtx, styling.Spoiler(text))
}

func (c *WotoContainer) SendBoldText(text string) (tg.UpdatesClass, error) {
	return c.GetSendBuilder().StyledText(gCtx, styling.Bold(text))
}

func (c *WotoContainer) SendItalicText(text string) (tg.UpdatesClass, error) {
	return c.GetSendBuilder().StyledText(gCtx, styling.Italic(text))
}

func (c *WotoContainer) SendMonoText(text string) (tg.UpdatesClass, error) {
	return c.GetSendBuilder().StyledText(gCtx, styling.Code(text))
}

func (c *WotoContainer) SendCodeText(text string) (tg.UpdatesClass, error) {
	return c.GetSendBuilder().StyledText(gCtx, styling.Code(text))
}

func (c *WotoContainer) SendStyledText(s wotoStyle.WStyle) (tg.UpdatesClass, error) {
	return c.GetSendBuilder().StyledText(gCtx, s.GetStylingArray()...)
}

func (c *WotoContainer) SendTextf(format string, args ...interface{}) (tg.UpdatesClass, error) {
	return c.GetSendBuilder().Textf(gCtx, format, args...)
}

func (c *WotoContainer) EditText(text string) (tg.UpdatesClass, error) {
	return c.GetEditBuilder().Text(gCtx, text)
}

func (c *WotoContainer) EditStrikeText(text string) (tg.UpdatesClass, error) {
	return c.GetEditBuilder().StyledText(gCtx, styling.Strike(text))
}

func (c *WotoContainer) EditBoldText(text string) (tg.UpdatesClass, error) {
	return c.GetEditBuilder().StyledText(gCtx, styling.Bold(text))
}

func (c *WotoContainer) EditItalicText(text string) (tg.UpdatesClass, error) {
	return c.GetEditBuilder().StyledText(gCtx, styling.Italic(text))
}

func (c *WotoContainer) EditMonoText(text string) (tg.UpdatesClass, error) {
	return c.GetEditBuilder().StyledText(gCtx, styling.Code(text))
}

func (c *WotoContainer) EditCodeText(text string) (tg.UpdatesClass, error) {
	return c.GetEditBuilder().StyledText(gCtx, styling.Code(text))
}

func (c *WotoContainer) EditStyledText(s wotoStyle.WStyle) (tg.UpdatesClass, error) {
	return c.GetEditBuilder().StyledText(gCtx, s.GetStylingArray()...)
}

func (c *WotoContainer) EditTextf(format string, args ...interface{}) (tg.UpdatesClass, error) {
	return c.GetEditBuilder().Textf(gCtx, format, args...)
}

func (c *WotoContainer) GetReplyBuilder() *message.Builder {
	return wg.SenderHelper.Reply(*c.Entities, c.GetAnswerable())
}

func (c *WotoContainer) GetSendBuilder() *message.RequestBuilder {
	return wg.SenderHelper.Answer(*c.Entities, c.GetAnswerable())
}

func (c *WotoContainer) GetEditBuilder() *message.EditMessageBuilder {
	return wg.SenderHelper.Answer(*c.Entities, c.GetAnswerable()).Edit(c.Message.ID)
}

func (c *WotoContainer) GetPrefixes() []rune {
	return wotoConfig.GetPrefixes()
}

func (c *WotoContainer) GetMessageText() string {
	if c.Message != nil {
		return c.Message.Message
	}

	return ""
}

func (c *WotoContainer) Args() *argparser.EventArgs {
	args, _ := argparser.ParseArg(c.GetMessageText(), c.GetPrefixes())
	return args
}

func (c *WotoContainer) GetClient() *telegram.Client {
	return wg.Client
}

func (c *WotoContainer) ResolveUsername(username string) (*tg.ContactsResolvedPeer, error) {
	return wg.API.ContactsResolveUsername(gCtx, username)
}

func (c *WotoContainer) ResolveUsernameToUser(username string) *tg.User {
	contacts, err := wg.API.ContactsResolveUsername(gCtx, username)
	if err != nil || contacts == nil {
		return nil
	}

	if len(contacts.Users) > 0 {
		u, ok := contacts.Users[0].AsNotEmpty()
		if u != nil && ok {
			if tgUtils.SaveTgUser != nil {
				// try to cache the user
				go tgUtils.SaveTgUser(u)
			}

			return u
		}
	}

	return nil
}

//---------------------------------------------------------
