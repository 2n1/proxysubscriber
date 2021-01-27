package forms

type UpdateAuthForm struct {
	Email string `form:"email" binding:"required"`
	Password string `form:"password"  binding:"required"`
	NPassword string `form:"npassword"  binding:"required"`
	RPassword string `form:"rpassword"  binding:"required"`
}

type LoginForm struct {
	Email string `form:"email" binding:"required"`
	Password string `form:"password"  binding:"required"`

}