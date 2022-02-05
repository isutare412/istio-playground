package user

import "context"

type Rest interface {
	GetUser(ctx context.Context, name string) (*User, error)
}
