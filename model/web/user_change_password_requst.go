package web

type UserChangePasswordRequest struct {
	Password      string `validate:"required,min=8,max=100" json:"password"`
	NewPassword   string `validate:"required,min=8,max=100" json:"new_password"`
	ReNewPassword string `validate:"required,min=8,max=100" json:"re_new_password"`
}
