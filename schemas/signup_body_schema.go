package schemas

type SignupBody struct {
	Name            string `form:"name" json:"name" validate:"required"`
	Email           string `form:"email" json:"email" validate:"required,email"`
	Password        string `form:"password" json:"password" validate:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" validate:"required,eqfield=Password"`
	DateOfBirth     string `form:"dateOfBirth" json:"dateOfBirth" validate:"required,datetime"`
	FirstName       string `form:"firstName" json:"firstName" validate:"required"`
	LastName        string `form:"lastName" json:"lastName" validate:"required"`
}
