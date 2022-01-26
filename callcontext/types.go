package callcontext

import (
	"context"
	"foodHelper/users"
)

type CallContext struct {
	context.Context
	user users.User
}
