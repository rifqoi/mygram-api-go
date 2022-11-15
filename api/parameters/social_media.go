package parameters

type SocialMediaCreate struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required,url"`
}

type SocialMediaUpdate struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url" validate:"url"`
}
