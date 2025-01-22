package auth

import (
	"context"

	"github.com/shelllbyyyyy/belajar-api-go/internal/exception"
)

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u UserUseCase) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	result, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		err = exception.ErrNotFound
		return nil, err
	}

	return result, nil
}

func (u UserUseCase) FindUserById(ctx context.Context, id string) (*User, error) {
	result, err := u.repo.FindById(ctx, id)
	if err != nil {
		err = exception.ErrNotFound
		return nil, err
	}

	return result, nil
}

func (u UserUseCase) Update(ctx context.Context, user *User, req updateUserSchema) (bool, error) {
	if req.Password != nil {
		if err := user.comparePassword(*req.Password); err != nil {
			return false, err
		}

		user.Password = *req.Password
		if err := user.encryptPassword(int(10)); err != nil {
			return false, err
		}
	}

	payload := &updateUserPayload{
		Id: user.Id,
		Username: req.Username,
		Email: req.Email,
		Password: func() *string {
			if req.Password != nil {
				return &user.Password
			}
			return nil
		}(),
	}

	return u.repo.Update(ctx, payload)
}