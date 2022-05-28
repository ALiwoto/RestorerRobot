package wotoAuth

import (
	"context"

	"github.com/ALiwoto/wotoub/wotoub/core/utils/logging"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoConfig"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoEntry"
	wv "github.com/ALiwoto/wotoub/wotoub/core/wotoValues"
	wg "github.com/ALiwoto/wotoub/wotoub/core/wotoValues/wotoGlobals"
	"github.com/ALiwoto/wotoub/wotoub/wotoActions"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	tgMessage "github.com/gotd/td/telegram/message"
	updhook "github.com/gotd/td/telegram/updates/hook"
	"github.com/gotd/td/tg"
)

func GetNewTerminalAuth(phoneNum string) auth.UserAuthenticator {
	return &TermAuth{
		phone: phoneNum,
	}
}

func AuthorizeClient(phone string, updateHandler wv.WotoUpdateHandler) error {
	err := wotoConfig.PrepareVariables()
	if err != nil {
		return err
	}

	ctx := context.Background()
	// Setting up authentication flow helper based on terminal auth.
	flow := auth.NewFlow(
		GetNewTerminalAuth(phone),
		auth.SendCodeOptions{},
	)

	//wv.UpdateManager = updates.New(updates.Config{
	//	Handler: updateHandler,
	//	//Logger: log.Named("gaps"),
	//})

	dTmp := tg.NewUpdateDispatcher()
	dispatcher := &dTmp

	wDispatcher := &wotoUpdateHandler{
		cachingHandler: updateHandler,
		realDispather:  dispatcher,
	}

	client, err := telegram.ClientFromEnvironment(telegram.Options{
		//Logger: log,
		//UpdateHandler: wv.UpdateManager,
		UpdateHandler: wDispatcher,
		Middlewares: []telegram.Middleware{
			//updhook.UpdateHook(uHandler),
			updhook.UpdateHook(wDispatcher.Handle),
		},
	})
	if err != nil {
		return err
	}

	wg.Client = client
	wg.SenderHelper = tgMessage.NewSender(client.API())

	wotoEntry.LoadAllHandlers(dispatcher)

	return client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return err
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
