package wotoGlobals

import (
	"context"
	"errors"

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

var (
	ErrPeerIdInvalid = errors.New("[400 PEER_ID_INVALID] - The peer id being used is invalid or not known yet. Make sure you meet the peer before interacting with it")
)
