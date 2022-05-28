package tgUtils

import (
	"context"

	"github.com/gotd/td/tg"
)

var (
	GetUserFromIdHelper func(id int64) (*tg.InputUser, error)
	SaveTgUser          func(u *tg.User) error
	gCtx                = context.Background()
)
