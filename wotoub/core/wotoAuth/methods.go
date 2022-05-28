package wotoAuth

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	terminal "golang.org/x/term"
)

//---------------------------------------------------------

func (c NoSignUp) SignUp(ctx context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, errors.New("not implemented")
}

func (c NoSignUp) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	return &auth.SignUpRequired{TermsOfService: tos}
}

//---------------------------------------------------------

func (t *TermAuth) Phone(_ context.Context) (string, error) {
	return t.phone, nil
}

func (t *TermAuth) Password(_ context.Context) (string, error) {
	fmt.Print("Enter 2FA password: ")
	bytePwd, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytePwd)), nil
}

func (t *TermAuth) Code(_ context.Context, _ *tg.AuthSentCode) (string, error) {
	fmt.Print("Enter code: ")
	code, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(code), nil
}

//---------------------------------------------------------

func (w *wotoUpdateHandler) Handle(ctx context.Context, u tg.UpdatesClass) error {
	if w.cachingHandler != nil {
		w.cachingHandler(ctx, u)
	}

	if w.realDispather != nil {
		return w.realDispather.Handle(ctx, u)
	}

	return nil
}
