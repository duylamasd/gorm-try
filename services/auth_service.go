package services

import (
	"context"
	"os"
	"time"

	"github.com/duylamasd/gorm-try/interfaces"
	"github.com/duylamasd/gorm-try/schemas"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	usersRepository interfaces.UsersRepository
}

func NewAuthService(usersRepository interfaces.UsersRepository) interfaces.AuthService {
	return &AuthService{usersRepository: usersRepository}
}

func (service *AuthService) generateAccessToken(uid string) (string, error) {
	secret := os.Getenv("JWT_ACCESS_SECRET")

	expiresAt := time.Now().Add(time.Minute * 15)

	claims := schemas.JwtAccessTokenClaims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *AuthService) generateRefreshToken(uid string) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")

	expiresAt := time.Now().Add(time.Hour * 10)

	claims := schemas.JwtRefreshTokenClaims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *AuthService) SignUp(ctx context.Context, payload *schemas.SignupBody) (*schemas.UserTokenInfo, error) {
	user, err := service.usersRepository.Create(payload)
	if err != nil {
		return nil, err
	}

	userID := user.ID.String()

	accessToken, err := service.generateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := service.generateRefreshToken(userID)
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
