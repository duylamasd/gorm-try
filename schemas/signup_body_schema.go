package schemas

type SignupBody struct {
	Name            string `form:"name" json:"name" binding:"required"`
	Email           string `form:"email" json:"email" binding:"required,email"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required,eqfield=Password"`
	DateOfBirth     string `form:"dateOfBirth" json:"dateOfBirth" binding:"required,datetime=2006-01-02"`
	FirstName       string `form:"firstName" json:"firstName" binding:"required"`
	LastName        string `form:"lastName" json:"lastName" binding:"required"`
}
