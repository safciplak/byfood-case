package users

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}
