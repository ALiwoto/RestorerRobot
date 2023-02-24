package wotoAuth

import (
	"context"
	"encoding/hex"
	"net/url"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/logging"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry"
	wv "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues"
	wg "github.com/AnimeKaizoku/RestorerRobot/src/core/wotoValues/wotoGlobals"
	"github.com/AnimeKaizoku/RestorerRobot/src/wotoActions"
	"github.com/go-faster/errors"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/dcs"
	tgMessage "github.com/gotd/td/telegram/message"
	updHook "github.com/gotd/td/telegram/updates/hook"
	"github.com/gotd/td/tg"
)

func getProxy() dcs.Resolver {
	proxyStr := wotoConfig.GetProxy()
	if proxyStr == "" {
		return nil
	}

	urlValue, err := url.Parse(proxyStr)
	if err != nil {
		return nil
	}

	queries := urlValue.Query()
	if len(queries) < 3 {
		return nil
	}

	pAddr := queries["server"][0] + ":" + queries["port"][0]
	bSecret, err := hex.DecodeString(queries["secret"][0])
	if err != nil {
		logging.Error("failed to decode mtproto proxy secret: " + err.Error())
		return nil
	}

	proxyValue, err := dcs.MTProxy(pAddr, bSecret, dcs.MTProxyOptions{})
	if err != nil {
		logging.Error("failed to create mtproto proxy: " + err.Error())
		return nil
	}

	return proxyValue
}

func AuthorizeClient(updateHandler wv.WotoUpdateHandler) error {
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

	client := telegram.NewClient(wotoConfig.GetAppId(), wotoConfig.GetAppHash(), telegram.Options{
		//Logger: log,
		//UpdateHandler: wv.UpdateManager,
		UpdateHandler: wDispatcher,
		Middlewares: []telegram.Middleware{
			//updHook.UpdateHook(uHandler),
			updHook.UpdateHook(wDispatcher.Handle),
		},
		Device: telegram.DeviceConfig{
			DeviceModel:   "Android 12 Snow Cone",
			SystemVersion: "12.0.13089",
			AppVersion:    "1.0.3",
		},
		Resolver: GetProxy(),
		SessionStorage: &telegram.FileSessionStorage{
			Path: wotoConfig.GetSessionPath(),
		},
	})

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
