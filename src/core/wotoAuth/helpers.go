package wotoAuth

import (
	"context"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/logging"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry"
	wv "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues"
	wg "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/AnimeKaizoku/RestorerRobot/src/wotoActions"
	"github.com/go-faster/errors"
	"github.com/gotd/td/telegram"
	tgMessage "github.com/gotd/td/telegram/message"
	updHook "github.com/gotd/td/telegram/updates/hook"
	"github.com/gotd/td/tg"
)

func AuthorizeClient(updateHandler wv.WotoUpdateHandler) error {
	err := wotoConfig.PrepareVariables()
	if err != nil {
		return err
	}

	ctx := context.Background()

	//wv.UpdateManager = updates.New(updates.Config{
	//	Handler: updateHandler,
	//	//Logger: log.Named("gaps"),
	//})

	dTmp := tg.NewUpdateDispatcher()
	dispatcher := &dTmp

	wDispatcher := &wotoUpdateHandler{
		cachingHandler: updateHandler,
		realDispatcher: dispatcher,
	}

	client, err := telegram.ClientFromEnvironment(telegram.Options{
		//Logger: log,
		//UpdateHandler: wv.UpdateManager,
		UpdateHandler: wDispatcher,
		Middlewares: []telegram.Middleware{
			//updHook.UpdateHook(uHandler),
			updHook.UpdateHook(wDispatcher.Handle),
		},
	})
	if err != nil {
		return err
	}

	wg.Client = client
	wg.SenderHelper = tgMessage.NewSender(client.API())

	wotoEntry.LoadAllHandlers(dispatcher)

	return client.Run(ctx, func(ctx context.Context) error {
		status, err := client.Auth().Status(ctx)
		if err != nil {
			return errors.Wrap(err, "auth status")
		}

		if !status.Authorized {
			if _, err := client.Auth().Bot(ctx, wotoConfig.GetBotToken()); err != nil {
				return errors.Wrap(err, "login")
			}
		}

		wg.Self, err = client.Self(ctx)
		if err != nil {
			return err
		}

		wg.API = client.API()

		wotoActions.LoadAllHandlers()

		// Notify update manager about authentication.
		//if err := wv.UpdateManager.Auth(ctx, wv.API, wv.Self.ID, wv.Self.Bot, true); err != nil {
		//	return err
		//}
		//defer func() {
		//	_ = wv.UpdateManager.Logout()
		//}()

		logging.Info("Authorized as ", wg.Self.FirstName)

		<-ctx.Done()

		return ctx.Err()
	})
}
