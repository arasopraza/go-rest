package user

import (
	"context"
	"fmt"
)

type usecase struct {
	userRepo Repository
}

func NewUsecase(userRepo Repository) Usecase {
	return &usecase{
		userRepo: userRepo,
	}
}

func (u *usecase) CreateUser(ctx context.Context, data User) (id string, createdAt string, err error) {

	id, createdAt, err = u.userRepo.CreateUser(data)

	if err != nil {
		err = fmt.Errorf("failed to create user: %s", err)
		return
	}

	return

}
