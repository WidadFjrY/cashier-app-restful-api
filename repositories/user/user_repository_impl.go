package repositories

import (
	"context"
	"database/sql"
	"github.com/widadfjry/cashier-app/exception"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.UserUpdateOrCreate {
	SQL := "INSERT INTO users(username, password, name, email, role, addedBy) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Name, user.Email, user.Role, user.AddedBy)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)

	userCreate := domain.UserUpdateOrCreate{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}

	return userCreate
}

func (repository *UserRepositoryImpl) Find(ctx context.Context, tx *sql.Tx, userId int, username string) domain.User {
	var user domain.User
	if username != "" {
		SQL := "SELECT id, username, name, email, role, addedBy, createdAt, updatedAt FROM users WHERE username = ?"
		rows, err := tx.QueryContext(ctx, SQL, username)
		helper.PanicIfError(err)
		defer rows.Close()

		if rows.Next() {
			err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Role, &user.AddedBy, &user.CreatedAt, &user.UpdatedAt)
			helper.PanicIfError(err)
		} else {
			panic(exception.NotFoundErrors{Error: "username not found"})
		}
	} else if userId > 0 {
		SQL := "SELECT id, username, name, email, role, addedBy, createdAt, updatedAt FROM users WHERE id = ?"
		rows, err := tx.QueryContext(ctx, SQL, userId)
		helper.PanicIfError(err)
		defer rows.Close()

		if rows.Next() {
			err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Role, &user.AddedBy, &user.CreatedAt, &user.UpdatedAt)
			helper.PanicIfError(err)
		} else {
			panic(exception.NotFoundErrors{Error: "id not found"})
		}
	}

	return user
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, offset int) []domain.User {
	SQL := "SELECT id, username, name, email, role, addedBy, createdAt, updatedAt FROM users  LIMIT 10 OFFSET ?"
	rows, err := tx.QueryContext(ctx, SQL, offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Role, &user.AddedBy, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.UserUpdateOrCreate, userId int, username string) domain.UserUpdateOrCreate {
	if username != "" {
		SQL := "UPDATE users SET email = ?, name = ? WHERE username = ?"
		_, err := tx.ExecContext(ctx, SQL, user.Email, user.Name, username)
		helper.PanicIfError(err)
	} else if userId != 0 {
		SQL := "UPDATE users SET email = ?, name = ? WHERE id = ?"
		_, err := tx.ExecContext(ctx, SQL, user.Email, user.Name, userId)
		helper.PanicIfError(err)
	}

	return user
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User, username string, email string) {
	if username != "" {
		SQL := "UPDATE users SET password = ? WHERE username = ?"
		_, err := tx.ExecContext(ctx, SQL, user.Password, username)
		helper.PanicIfError(err)
	} else if email != "" {
		SQL := "UPDATE users SET password = ? WHERE email = ?"
		_, err := tx.ExecContext(ctx, SQL, user.Password, email)
		helper.PanicIfError(err)
	}
}

func (repository *UserRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, userId int) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) Count(ctx context.Context, tx *sql.Tx) int {
	SQL := "SELECT COUNT(*) AS total_super_admin FROM users WHERE role = 'super_admin'"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var totalSuperAdmin int

	if rows.Next() {
		err := rows.Scan(&totalSuperAdmin)
		helper.PanicIfError(err)
	}

	return totalSuperAdmin
}

func (repository *UserRepositoryImpl) UsernameCheck(ctx context.Context, tx *sql.Tx) []domain.UserUsername {
	SQL := "SELECT username FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var usernames []domain.UserUsername

	for rows.Next() {
		username := domain.UserUsername{}
		err := rows.Scan(&username.Username)
		helper.PanicIfError(err)

		usernames = append(usernames, username)
	}

	return usernames
}

func (repository *UserRepositoryImpl) EmailCheck(ctx context.Context, tx *sql.Tx) []domain.UserEmail {
	SQL := "SELECT email FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var emails []domain.UserEmail

	for rows.Next() {
		email := domain.UserEmail{}
		err := rows.Scan(&email.Email)
		helper.PanicIfError(err)

		emails = append(emails, email)
	}

	return emails
}

func (repository *UserRepositoryImpl) CountUsers(ctx context.Context, tx *sql.Tx) int {
	SQL := "SELECT COUNT(*) AS total_users FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var totalUser int
	for rows.Next() {
		err := rows.Scan(&totalUser)
		helper.PanicIfError(err)
	}

	return totalUser
}

func (repository *UserRepositoryImpl) FindPassword(ctx context.Context, tx *sql.Tx, username string) string {
	SQL := "SELECT password FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	var userPassword string
	if rows.Next() {
		err := rows.Scan(&userPassword)
		helper.PanicIfError(err)
	} else {
		panic(exception.NewBadRequestErrors("username or password wrong"))
	}

	return userPassword
}

func (repository *UserRepositoryImpl) FindIdByUsername(ctx context.Context, tx *sql.Tx, username string) int {
	SQL := "SELECT id FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		helper.PanicIfError(err)
	} else {
		panic(exception.NewNotFoundErrors("username not found"))
	}

	return id

}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) string {
	SQL := "SELECT name FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()

	var name string
	if rows.Next() {
		err := rows.Scan(&name)
		helper.PanicIfError(err)
	}

	return name
}

func (repository *UserRepositoryImpl) Logout(ctx context.Context, tx *sql.Tx, token domain.BlockedToken) {
	SQL := "INSERT INTO blockedusers(token) VALUES (?)"
	_, err := tx.ExecContext(ctx, SQL, token.Token)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) SaveOTP(ctx context.Context, tx *sql.Tx, otp domain.OTP) {
	SQL := "INSERT INTO otp (email, otp) VALUES (?, ?)"
	_, err := tx.ExecContext(ctx, SQL, otp.Email, otp.OTP)
	if err != nil {
		panic(exception.NewBadRequestErrors("email not registered"))
	}
}

func (repository *UserRepositoryImpl) FindEmailByOTP(ctx context.Context, tx *sql.Tx, otp string) string {
	SQL := "SELECT email FROM otp WHERE otp = ?"
	rows, err := tx.QueryContext(ctx, SQL, otp)
	helper.PanicIfError(err)
	defer rows.Close()

	var email string
	if rows.Next() {
		err := rows.Scan(&email)
		helper.PanicIfError(err)
	}

	return email
}
