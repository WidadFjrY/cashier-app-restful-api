package web

type UserUpdateRequest struct {
	Name  string `validate:"required,min=1,max=100" json:"name"`
	Email string `validate:"required,email,min=1,max=100" json:"email"`
}
