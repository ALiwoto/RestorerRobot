package tgUtils

import (
	"context"

	"github.com/ALiwoto/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/gotd/td/tg"
)

var (
	GetUserFromIdHelper func(id int64) (*tg.InputUser, error)
	GetInputPeerInfo    func(id int64) (*wotoGlobals.PeerInfo, error)
	SaveTgUser          func(u *tg.User) error
	gCtx                = context.Background()
)
