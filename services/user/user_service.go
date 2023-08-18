package services

import (
	"context"
	"github.com/widadfjry/cashier-app/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserUpdateOrCreateResponse
	AddUser(ctx context.Context, request web.UserCreateRequest, isAdmin bool, username string) web.UserUpdateOrCreateResponse
	Login(ctx context.Context, request web.UserLoginRequest) string
	FindById(ctx context.Context, userId int) web.UserResponses
	FindAll(ctx context.Context, page int) ([]web.UserResponses, web.Pages)
	UpdateById(ctx context.Context, request web.UserUpdateRequest, userId int, username string, isAdmin bool) web.UserUpdateOrCreateResponse
	UpdateByUsername(ctx context.Context, request web.UserUpdateRequest, username string) web.UserUpdateOrCreateResponse
	DeleteById(ctx context.Context, userId int, isAdmin bool)
	UpdatePasswordByUsername(ctx context.Context, request web.UserChangePasswordRequest, username string)
	Logout(ctx context.Context, token string)
	Verification(ctx context.Context, request web.UserVerificationRequest) string
	NewPassword(ctx context.Context, email string, request web.UserNewPasswordRequest) error
}
