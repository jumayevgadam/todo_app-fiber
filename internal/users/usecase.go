package users

import (
	"context"

	userModel "github.com/jumayevgadam/todo_app-fiber/internal/models/users"
)

// Service interface is
type Service interface {
	SignUp(ctx context.Context, reqDTO *userModel.SignUpReq) (int64, error)
}
