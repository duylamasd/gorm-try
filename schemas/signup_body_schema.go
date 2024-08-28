package schemas

type SignupBody struct {
	Name            string `form:"name" json:"name" validate:"required" binding:"required"`
	Email           string `form:"email" json:"email" validate:"required,email" binding:"required,email"`
	Password        string `form:"password" json:"password" validate:"required" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" validate:"required,eqfield=Password" binding:"required,eqfield=Password"`
	DateOfBirth     string `form:"dateOfBirth" json:"dateOfBirth" validate:"required,datetime" binding:"required,datetime"`
	FirstName       string `form:"firstName" json:"firstName" validate:"required" binding:"required"`
	LastName        string `form:"lastName" json:"lastName" validate:"required" binding:"required"`
}
