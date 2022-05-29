package wotoAuth

import (
	"context"

	"github.com/gotd/td/tg"
)

//---------------------------------------------------------

func (w *wotoUpdateHandler) Handle(ctx context.Context, u tg.UpdatesClass) error {
	if w.cachingHandler != nil {
		w.cachingHandler(ctx, u)
	}

	if w.realDispatcher != nil {
		return w.realDispatcher.Handle(ctx, u)
	}

	return nil
}
