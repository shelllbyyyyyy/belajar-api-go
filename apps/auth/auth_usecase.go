package auth

import (
	"context"

	"github.com/shelllbyyyyy/belajar-api-go/util"
)

type AuthUseCase struct {
	repo UserRepository
}

func NewAuthUseCase(repo UserRepository) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

func (u AuthUseCase) CreateUser(ctx context.Context, req registerUserSchema) (*string, error) {
	user, err := newUser(req)
	if err != nil {
		return nil, err
	}

	err = user.validate()
	if err != nil {
		return nil, err
	}

	err = user.encryptPassword(10)
	if err != nil {
		return nil, err
	}

	result, err := u.repo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u AuthUseCase) ValidateUserCredentials(ctx context.Context, user *User, password string) (*token, error) {
	err := user.comparePassword(password); if err != nil {
		return nil, err
	}

	accessToken, err := user.generateToken(15); if err != nil {
		return nil, err
	}

	refreshToken, err := user.generateToken(60 * 7 * 24); if err != nil {
		return nil, err
	}

	return &token{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
 }

 func (u AuthUseCase) Refresh(ctx context.Context, req tokenSchema) (string, error) {
	token, err := util.GenerateToken(req.Id, 15)
	if err != nil {
		return "", err
	}
	
	return token, nil	
}