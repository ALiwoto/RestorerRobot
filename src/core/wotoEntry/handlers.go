package wotoEntry

import (
	"context"

	"github.com/ALiwoto/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/gotd/td/tg"
)

func NewMessageHandler(ctx context.Context, entities tg.Entities, u *tg.UpdateNewMessage) error {
	if u.Message != nil {
		container := &entryManager.WotoContainer{
			TelegramClient:   wotoGlobals.Client,
			OriginMessage:    u.Message,
			Update:           u,
			UpdateNewMessage: u,
			Entities:         &entities,
			Answerable:       u,
		}
		handleNewMessage(container)
	}
	return nil
}

func NewChannelMessageHandler(ctx context.Context, entities tg.Entities, u *tg.UpdateNewChannelMessage) error {
	if u.Message != nil {
		container := &entryManager.WotoContainer{
			TelegramClient:          wotoGlobals.Client,
			OriginMessage:           u.Message,
			Update:                  u,
			UpdateNewChannelMessage: u,
			Entities:                &entities,
			Answerable:              u,
		}
		handleNewMessage(container)
	}
	return nil
}

func NewScheduledMessageHandler(ctx context.Context, entities tg.Entities, u *tg.UpdateNewScheduledMessage) error {
	if u.Message != nil {
		container := &entryManager.WotoContainer{
			TelegramClient:            wotoGlobals.Client,
			OriginMessage:             u.Message,
			Update:                    u,
			UpdateNewScheduledMessage: u,
			Entities:                  &entities,
			Answerable:                u,
		}
		handleNewMessage(container)
	}
	return nil
}

func EditMessageHandler(ctx context.Context, entities tg.Entities, u *tg.UpdateEditMessage) error {
	if u.Message != nil {
		container := &entryManager.WotoContainer{
			TelegramClient:    wotoGlobals.Client,
			OriginMessage:     u.Message,
			Update:            u,
			UpdateEditMessage: u,
			Entities:          &entities,
			Answerable:        u,
		}
		handleNewMessage(container)
	}
	return nil
}
