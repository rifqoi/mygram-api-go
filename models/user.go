package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;notNull"`
	Email     string    `json:"email" gorm:"uniqueIndex;notNull"`
	Username  string    `json:"username" gorm:"uniqueIndex;notNull"`
	Password  string    `json:"password" gorm:"notNull"`
	Age       int       `json:"age" gorm:"notNull"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
