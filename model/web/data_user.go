package web

type GetDataUser struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	AddedBy   string `json:"added_by"`
	CratedAt  string `json:"crated_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateDataUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
