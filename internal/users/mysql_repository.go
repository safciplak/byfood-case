package users

import (
	"context"
	"errors"
	"github.com/byfood/byfood-core/app"
	"gorm.io/gorm"
)

func NewMysqlRepo(db *gorm.DB) *MysqlRepo {
	return &MysqlRepo{db: db}
}

type MysqlRepo struct {
	db *gorm.DB
}

func (r *MysqlRepo) GetUser(ctx context.Context, userID int64) (*User, error) {
	var u User
	if err := r.db.First(&u, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrorNotfound
		}
		return nil, app.WrapError(err)
	}

	return &u, nil
}

func (r *MysqlRepo) AddUser(ctx context.Context, u *User) error {
	if err := r.db.Create(u).Error; err != nil {
		return app.WrapError(err)
	}
	return nil
}

func (r *MysqlRepo) UpdateUser(ctx context.Context, u *User) error {
	if err := r.db.Updates(u).Error; err != nil {
		return app.WrapError(err)
	}
	return nil
}

func (r *MysqlRepo) DeleteUser(ctx context.Context, u *User) error {
	if err := r.db.Delete(u).Error; err != nil {
		return app.WrapError(err)
	}
	return nil
}

func (r *MysqlRepo) GetAllUsers(ctx context.Context) ([]User, error) {
	var u []User
	if err := r.db.Find(&u).Error; err != nil {
		return nil, app.WrapError(err)
	}
	return u, nil
}
