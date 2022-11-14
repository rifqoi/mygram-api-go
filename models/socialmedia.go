package models

import "time"

type SocialMedia struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"notNull"`
	SocialMediaUrl string `json:"social_media_url" gorm:"notNull"`
	UserID         int    `json:"user_id" gorm:"notNull"`

	// User disini untuk mereferensi UserID
	// https://stackoverflow.com/questions/73661393/has-one-vs-belongs-to-in-gorm
	User      User      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SocialMediaUserGet struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type SocialMediaGet struct {
	ID             int                `json:"id" gorm:"primaryKey"`
	Name           string             `json:"name" gorm:"notNull"`
	SocialMediaUrl string             `json:"social_media_url" gorm:"notNull"`
	UserID         int                `json:"user_id" gorm:"notNull"`
	User           SocialMediaUserGet `json:"user"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
