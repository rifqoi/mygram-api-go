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

type CommentUser struct {
	ID       int    `json:"id" gorm:"primaryKey;notNull"`
	Email    string `json:"email" gorm:"uniqueIndex;notNull"`
	Username string `json:"username" gorm:"uniqueIndex;notNull"`
}

type CommentPhoto struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type CommentGetAll struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Message string `json:"message" gorm:"notNull"`
	PhotoID int    `json:"photo_id" gorm:"notNull"`
	UserID  int    `json:"user_id" gorm:"notNull"`
	User    CommentUser
	Photo   CommentPhoto
}
