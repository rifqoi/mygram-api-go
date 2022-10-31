package parameters

type UserRegister struct {
	Age      int    `json:"age" validate:"required,gte=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,validatepassword"`
	Username string `json:"username" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,validatepassword"`
}

type UserUpdate struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}
