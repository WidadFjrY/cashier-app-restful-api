package web

type UserNewPasswordRequest struct {
	OTP           string `validate:"required,min=8, max=8" json:"otp"`
	NewPassword   string `validate:"required,min=8,max100" json:"new_password"`
	ReNewPassword string `validate:"required,min=8,max100" json:"re_new_password"`
}
