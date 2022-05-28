package wotoGlobals

import (
	"context"

	"github.com/gotd/td/telegram"
	tgMessage "github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/tg"
)

var (
	GCtx          = context.Background()
	Client        *telegram.Client
	API           *tg.Client
	Self          *tg.User
	UpdateManager *updates.Manager
	SenderHelper  *tgMessage.Sender
)
