package web

type UserVerificationRequest struct {
	Email string `validate:"required,min=1,max=100,email" json:"email"`
}
