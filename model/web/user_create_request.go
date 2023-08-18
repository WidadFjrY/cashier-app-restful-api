package web

type UserCreateRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=8,max=100" json:"password"`
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Email    string `validate:"required,email,min=1,max=100" json:"email"`
	Role     string `validate:"required,min=1,max=20" json:"role"`
	AddedBy  string `validate:"max=100" json:"added_by"`
}
