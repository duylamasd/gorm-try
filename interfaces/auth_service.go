package interfaces

import (
	"context"

	"github.com/duylamasd/gorm-try/schemas"
)

type AuthService interface {
	SignUp(ctx context.Context, payload *schemas.SignupBody) (*schemas.UserTokenInfo, error)
}
