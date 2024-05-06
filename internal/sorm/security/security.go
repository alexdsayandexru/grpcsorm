package security

import "context"

type ISecurity interface {
	Auth(ctx *context.Context) (bool, error)
}
