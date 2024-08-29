package schemas

import "github.com/golang-jwt/jwt/v5"

type JwtRefreshTokenClaims struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}
