package models

import "time"

type Comment struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Message   string    `json:"message" gorm:"notNull"`
	PhotoID   int       `json:"photo_id" gorm:"notNull"`
	UserID    int       `json:"user_id" gorm:"notNull"`
	User      User      `json:"-"`
	Photo     Photo     `json:"-"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"-"`
}

type CommentUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentPhoto struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type CommentGetAll struct {
	ID        int          `json:"id" gorm:"primaryKey"`
	Message   string       `json:"message" gorm:"notNull"`
	PhotoID   int          `json:"photo_id" gorm:"notNull"`
	UserID    int          `json:"user_id" gorm:"notNull"`
	CreatedAt time.Time    `json:"create_at"`
	UpdatedAt time.Time    `json:"-"`
	User      CommentUser  `json:"-"`
	Photo     CommentPhoto `json:"-"`
}

type CommentUpdate struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"-"`
}
