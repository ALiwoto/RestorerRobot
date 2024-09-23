package wotoAuth

import (
	wv "github.com/ALiwoto/RestorerRobot/src/core/wotoValues"
	"github.com/gotd/td/tg"
)

type wotoUpdateHandler struct {
	cachingHandler wv.WotoUpdateHandler
	realDispatcher *tg.UpdateDispatcher
}
