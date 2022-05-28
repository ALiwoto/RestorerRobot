package wotoAuth

import (
	wv "github.com/ALiwoto/wotoub/wotoub/core/wotoValues"
	"github.com/gotd/td/tg"
)

// noSignUp can be embedded to prevent signing up.
type NoSignUp struct{}

// termAuth implements authentication via terminal.
type TermAuth struct {
	NoSignUp

	phone string
}

type wotoUpdateHandler struct {
	cachingHandler wv.WotoUpdateHandler
	realDispather  *tg.UpdateDispatcher
}
