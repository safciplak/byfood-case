package users

import (
	"context"
	"github.com/byfood/byfood-core/app"
)

var (
	ErrorNotfound = app.BusinessError("users not found")
)

type usersService interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, userID int64) (*User, error)
	AddUser(ctx context.Context, u *User) error
	UpdateUser(ctx context.Context, u *User) error
	DeleteUser(ctx context.Context, u *User) error
}

type Service struct {
	Users usersService
}

func (s *Service) GetAllUsers(ctx context.Context) ([]User, error) {
	return s.Users.GetAllUsers(ctx)
}

func (s *Service) GetUser(ctx context.Context, userID int64) (*User, error) {
	return s.Users.GetUser(ctx, userID)
}

func (s *Service) AddUser(ctx context.Context, u *User) error {
	return s.Users.AddUser(ctx, u)
}

func (s *Service) UpdateUser(ctx context.Context, u *User) error {
	return s.Users.UpdateUser(ctx, u)
}

func (s *Service) DeleteUser(ctx context.Context, u *User) error {
	return s.Users.DeleteUser(ctx, u)
}
