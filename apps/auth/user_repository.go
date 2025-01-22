package auth

import "context"

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindById(ctx context.Context, id string) (*User, error)
	Register(ctx context.Context, model *User) (string, error)
	Update(ctx context.Context, payload *updateUserPayload) (bool, error)
}