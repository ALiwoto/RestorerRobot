package enteryManager

import (
	"context"
	"errors"
)

// error vars
var (
	ErrEndGroups      = errors.New("end of groups")
	ErrContinueGroups = errors.New("continue groups")
)

// internal variables
var (
	gCtx = context.Background()
)
