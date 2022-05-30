package entryManager

import (
	"context"
	"path"
	"strings"
	"sync"

	"github.com/ALiwoto/argparser/argparser"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/tgUtils"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoStyle"
	wg "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
)

//---------------------------------------------------------

func (m *EntryManager) IsEmpty() bool {
	return len(m.entryMap) == 0
}

func (m *EntryManager) GetEntry(name string) *entry {
	if m.entryMutex == nil {
		m.entryMutex = &sync.Mutex{}
	}
	m.entryMutex.Lock()
	e := m.entryMap[name]
	m.entryMutex.Unlock()
	return e
}

func (m *EntryManager) AddEntry(name string, e *entry) {
	if m.entryMutex == nil {
		m.entryMutex = &sync.Mutex{}
	}
	m.entryMutex.Lock()
	m.entryMap[name] = e
	m.entryMutex.Unlock()
}

func (m *EntryManager) AddHandlers(name string, h ...MessageHandler) *entry {
	e := &entry{
		enabled:    true,
		handlers:   h,
		UniqueName: name,
	}

	e.internalCondition = func(message *tg.Message) bool {
		return len(message.Message) > 1 && strings.HasPrefix(message.Message[1:], name)
	}

	m.AddEntry(name, e)
	return e
}

func (m *EntryManager) RemoveEntry(name string) {
	if m.entryMutex == nil {
		m.entryMutex = &sync.Mutex{}
	}
	m.entryMutex.Lock()
	delete(m.entryMap, name)
	m.entryMutex.Unlock()
}

func (m *EntryManager) AddTriggers(t ...rune) []rune {
	m.triggers = append(m.triggers, t...)
	return m.triggers
}

func (m *EntryManager) GetTriggers() []rune {
	return m.triggers
}

func (m *EntryManager) SetTriggers(t []rune) {
	m.triggers = t
}

func (m *EntryManager) ShouldRevoke(message string) bool {
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

func (m *EntryManager) Revoke(container *WotoContainer) (next bool) {
	message := container.Message
	if m.ShouldRevoke(message.Message) {
		m.entryMutex.Lock()
		for _, entry := range m.entryMap {
			if entry.IsEnabled() && entry.ShouldRun(message) {
				next = entry.RunHandlers(container)
				if !next {
					break
				}
			}
		}
		m.entryMutex.Unlock()
		return false
	}

	return true
}

//---------------------------------------------------------

func (e *entry) IsEnabled() bool {
	return e.enabled
}

func (e *entry) Disable() {
	e.enabled = false
}

func (e *entry) Enable() {
	e.enabled = true
}

func (e *entry) IsEmpty() bool {
	return len(e.handlers) == 0
}

func (e *entry) AddHandlers(h ...MessageHandler) {
	e.handlers = append(e.handlers, h...)
}

func (e *entry) GetHandlers() []MessageHandler {
	return e.handlers
}

func (e *entry) SetHandlers(h []MessageHandler) {
	e.handlers = h
}

func (e *entry) ShouldRun(message *tg.Message) bool {
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

func (e *entry) RunHandlers(container *WotoContainer) (next bool) {
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

func (c *WotoContainer) ReplyError(description string, err error) (tg.UpdatesClass, error) {
	return c.ReplyStyledText(wotoStyle.GetBold(description).Normal(": " + err.Error()))
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

func (c *WotoContainer) GetSenderHelper() *message.Sender {
	return wg.SenderHelper
}

func (c *WotoContainer) UploadFileToChatByPath(filename string, opts *UploadDocumentOptions) error {
	uploader := uploader.NewUploader(wg.API).WithThreads(opts.Goroutines)
	sender := message.NewSender(wg.API).WithUploader(uploader)
	upload, err := uploader.FromPath(c.Ctx(), filename)
	if err != nil {
		return err
	}
	caption := opts.Caption

	builder := message.UploadedDocument(upload, caption.GetStylingArray()...)
	builder = builder.Filename(path.Base(filename))
	builder.ForceFile(true)

	inputTarget, err := tgUtils.GetInputPeerClass(opts.ChatID)
	if err != nil {
		return err
	}

	target := sender.To(inputTarget)
	if opts.ReplyToMessageId != 0 {
		_ = target.Reply(opts.ReplyToMessageId)
	}

	// Sending message with media.
	if _, err := target.Media(c.Ctx(), builder); err != nil {
		return err
	}

	return nil
}

func (c *WotoContainer) UploadFileToChatsByPath(filename string, opts *UploadDocumentToChatsOptions) error {
	uploader := uploader.NewUploader(wg.API).WithThreads(opts.Goroutines)
	sender := message.NewSender(wg.API).WithUploader(uploader)
	upload, err := uploader.FromPath(c.Ctx(), filename)
	if err != nil {
		return err
	}
	caption := opts.Caption

	builder := message.UploadedDocument(upload, caption.GetStylingArray()...)
	builder = builder.Filename(path.Base(filename))
	builder.ForceFile(true)

	for _, chatID := range opts.ChatIDs {
		inputTarget, err := tgUtils.GetInputPeerClass(chatID)
		if err != nil {
			return err
		}

		target := sender.To(inputTarget)

		// Sending message with media.
		if _, err := target.Media(c.Ctx(), builder); err != nil {
			return err
		}
	}

	return nil
}

func (c *WotoContainer) GetPrefixes() []rune {
	return wotoConfig.GetPrefixes()
}

// GetEffectiveUserID method returns the user-id of the effective user,
// (for example the person who sent the command/query, etc...)
func (c *WotoContainer) GetEffectiveUserID() int64 {
	switch {
	case c.Message != nil:
		return tgUtils.GetEffectiveUserIdFromMessage(c.Message)
		// TODO: add support for other cases...
	}
	return 0
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

func (c *WotoContainer) GetArgs() (*argparser.EventArgs, error) {
	return argparser.ParseArg(c.GetMessageText(), c.GetPrefixes())
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
