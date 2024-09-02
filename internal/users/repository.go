package users

import (
	"context"

	userModel "github.com/jumayevgadam/todo_app-fiber/internal/models/users"
)

// Repository is
type Repository interface {
	SignUp(ctx context.Context, resDAO *userModel.SignUpRes) (int64, error)
}
