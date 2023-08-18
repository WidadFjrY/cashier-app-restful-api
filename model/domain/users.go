package domain

type User struct {
	Id        int
	Username  string
	Password  string
	Name      string
	Email     string
	Role      string
	AddedBy   string
	CreatedAt string
	UpdatedAt string
}

type UserUpdateOrCreate struct {
	Id       int
	Username string
	Name     string
	Email    string
}

type UserUsername struct {
	Username string
}

type UserEmail struct {
	Email string
}
