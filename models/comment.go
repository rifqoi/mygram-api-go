package models

import "time"

type Comment struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Message   string    `json:"message" gorm:"notNull"`
	PhotoID   int       `json:"photo_id" gorm:"notNull"`
	UserID    int       `json:"user_id" gorm:"notNull"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
