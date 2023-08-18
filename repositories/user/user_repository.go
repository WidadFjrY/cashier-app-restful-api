package repositories

import (
	"context"
	"database/sql"
	"github.com/widadfjry/cashier-app/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.UserUpdateOrCreate
	Find(ctx context.Context, tx *sql.Tx, userId int, username string) domain.User
	FindAll(ctx context.Context, tx *sql.Tx, offset int) []domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.UserUpdateOrCreate, userId int, username string) domain.UserUpdateOrCreate
	UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User, username string, email string)
	DeleteById(ctx context.Context, tx *sql.Tx, userId int)
	Count(ctx context.Context, tx *sql.Tx) int
	UsernameCheck(ctx context.Context, tx *sql.Tx) []domain.UserUsername
	EmailCheck(ctx context.Context, tx *sql.Tx) []domain.UserEmail
	CountUsers(ctx context.Context, tx *sql.Tx) int
	FindPassword(ctx context.Context, tx *sql.Tx, username string) string
	FindIdByUsername(ctx context.Context, tx *sql.Tx, username string) int
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) string
	Logout(ctx context.Context, tx *sql.Tx, token domain.BlockedToken)
	SaveOTP(ctx context.Context, tx *sql.Tx, otp domain.OTP)
	FindEmailByOTP(ctx context.Context, tx *sql.Tx, otp string) string
}
