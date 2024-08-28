package services

import (
	"context"

	"github.com/duylamasd/gorm-try/interfaces"
	"github.com/duylamasd/gorm-try/schemas"
)

type AuthService struct {
	usersRepository interfaces.UsersRepository
}

func NewAuthService(usersRepository interfaces.UsersRepository) interfaces.AuthService {
	return &AuthService{usersRepository: usersRepository}
}

func (service *AuthService) generateAccessToken(uid string) (string, error) {
	return "", nil
}

func (service *AuthService) generateRefreshToken(uid string) (string, error) {
	return "", nil
}

func (service *AuthService) SignUp(ctx context.Context, payload *schemas.SignupBody) (*schemas.UserTokenInfo, error) {
	accessToken, err := service.generateAccessToken("")
	if err != nil {
		return nil, err
	}

	refreshToken, err := service.generateRefreshToken("")
	if err != nil {
		return nil, err
	}

	tokenInfo := &schemas.UserTokenInfo{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}

	return tokenInfo, nil
}
