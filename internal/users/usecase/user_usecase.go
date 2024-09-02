package usecase

import (
	"context"
	"fmt"

	"github.com/jumayevgadam/todo_app-fiber/internal/database"
	userModel "github.com/jumayevgadam/todo_app-fiber/internal/models/users"
)

// UserService struct is
type UserService struct {
	repo database.DataStore
}

// NewUserService is
func NewUserService(repo database.DataStore) *UserService {
	return &UserService{repo: repo}
}

// SignUP is
func (us *UserService) SignUp(ctx context.Context, reqDTO *userModel.SignUpReq) (int64, error) {
	var (
		userID int64
		err    error
	)

	if err := us.repo.WithTransaction(ctx, func(db database.DataStore) error {
		userID, err = db.UsersRepo().SignUp(ctx, reqDTO.ToStorage())
		if err != nil {
			return fmt.Errorf("[db.UsersRepo][SignUp]: %w", err)
		}

		return nil
	}); err != nil {
		return -1, fmt.Errorf("[userService][SignUP]: %w", err)
	}

	return userID, nil
}
