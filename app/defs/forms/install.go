package forms

type InstallForm struct {
	Email string `form:"email" binding:"required"`
	Password string `form:"password"  binding:"required"`

}