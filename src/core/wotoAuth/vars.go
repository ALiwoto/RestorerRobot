package wotoAuth

import "github.com/gotd/td/telegram/dcs"

var (
	GetProxy func() dcs.Resolver = getProxy
)
