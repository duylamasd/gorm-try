package schemas

import "github.com/golang-jwt/jwt/v5"

type JwtAccessTokenClaims struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}
