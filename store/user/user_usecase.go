package user

import (
	"context"
)

type Usecase interface {
	CreateUser(ctx context.Context, data User) (id string, createdAt string, err error)
}
