package parameters

type Comment struct {
	Message string `json:"message" validate:"required"`
	PhotoID int    `json:"photo_id" validate:"required"`
}
