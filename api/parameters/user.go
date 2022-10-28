package parameters

import (
	"github.com/go-playground/validator/v10"
)

type UserRegister struct {
	Age      int    `json:"age" validate:"required,gte=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,validatepassword"`
	Username string `json:"username" validate:"required"`
}

func (u UserRegister) Validate() []string {
	validate := validator.New()
	validate.RegisterValidation("validatepassword", ValidatePassword)

	// addTranslation("validatepassword", passwordErrMsg, validate)

	err := validate.Struct(u)
	if err != nil {
		errs := TranslateError(err, validate)
		return errs
	}
	return nil
}
