package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator"
	"github.com/widadfjry/cashier-app/exception"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/domain"
	"github.com/widadfjry/cashier-app/model/web"
	repositories "github.com/widadfjry/cashier-app/repositories/user"
	"golang.org/x/crypto/bcrypt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

var otps []string

func NewUserService(userRepository repositories.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserUpdateOrCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles := []string{"admin", "super_admin"}
	isValidRole := true

	err = service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	countSuperAdmin := service.UserRepository.Count(ctx, tx)
	if countSuperAdmin >= 100 {
		panic(exception.NewBadRequestErrors("super admin already exist"))
	}

	usernames := service.UserRepository.UsernameCheck(ctx, tx)
	for _, username := range usernames {
		if username.Username == request.Username {
			panic(exception.NewBadRequestErrors("username already in use"))
		}
	}

	emails := service.UserRepository.EmailCheck(ctx, tx)
	for _, email := range emails {
		if email.Email == request.Email {
			panic(exception.NewBadRequestErrors("email already in use"))
		}
	}

	for _, role := range roles {
		if role == request.Role {
			isValidRole = true
			break
		} else {
			isValidRole = false
		}
	}

	if !isValidRole {
		panic(exception.NewBadRequestErrors("role must be admin or super admin"))
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicIfError(err)

	users := domain.User{
		Username: request.Username,
		Name:     request.Name,
		Email:    request.Email,
		Password: string(passwordHash),
		Role:     request.Role,
		AddedBy:  request.AddedBy,
	}

	userCreate := service.UserRepository.Save(ctx, tx, users)

	return helper.ToUserUpdateOrCreateResponse(userCreate)
}

func (service *UserServiceImpl) AddUser(ctx context.Context, request web.UserCreateRequest, isAdmin bool, username string) web.UserUpdateOrCreateResponse {
	if !isAdmin {
		panic(exception.NewBadRequestErrors("you're not a super admin"))
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles := []string{"admin", "super_admin"}
	isValidRole := true

	err = service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	countSuperAdmin := service.UserRepository.Count(ctx, tx)
	if countSuperAdmin >= 100 {
		panic(exception.NewBadRequestErrors("super admin already exist"))
	}

	usernames := service.UserRepository.UsernameCheck(ctx, tx)
	for _, username := range usernames {
		if username.Username == request.Username {
			panic(exception.NewBadRequestErrors("username already in use"))
		}
	}

	emails := service.UserRepository.EmailCheck(ctx, tx)
	for _, email := range emails {
		if email.Email == request.Email {
			panic(exception.NewBadRequestErrors("email already in use"))
		}
	}

	for _, role := range roles {
		if role == request.Role {
			isValidRole = true
			break
		} else {
			isValidRole = false
		}
	}

	if !isValidRole {
		panic(exception.NewBadRequestErrors("role must be admin or super admin"))
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicIfError(err)

	users := domain.User{
		Username: request.Username,
		Name:     request.Name,
		Email:    request.Email,
		Password: string(passwordHash),
		Role:     request.Role,
		AddedBy:  username,
	}

	userCreate := service.UserRepository.Save(ctx, tx, users)

	return helper.ToUserUpdateOrCreateResponse(userCreate)

}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	userPassword := service.UserRepository.FindPassword(ctx, tx, request.Username)
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(request.Password))
	if err != nil {
		panic(exception.NewBadRequestErrors("username or password wrong"))
	}

	token, err := helper.CreateToken(request.Username)
	helper.PanicIfError(err)

	return token
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponses {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.Find(ctx, tx, userId, "")
	return helper.ToUserResponse(users)
}

func (service *UserServiceImpl) FindAll(ctx context.Context, page int) ([]web.UserResponses, web.Pages) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	totalUsers := service.UserRepository.CountUsers(ctx, tx)

	totalItems := totalUsers
	totalPages := math.Ceil(float64(totalItems) / 10.0)

	if float64(page) > totalPages {
		panic(exception.NewNotFoundErrors("page not found"))
	}

	strPage := ""
	if page == 1 {
		strPage = "0"
	} else {
		strPage = strconv.Itoa(page - 1)
	}

	strPage = strPage + "0"

	offset, err := strconv.Atoi(strPage)
	helper.PanicIfError(err)

	users := service.UserRepository.FindAll(ctx, tx, offset)
	webPages := web.Pages{
		Page:       page,
		TotalPages: totalPages,
		TotalItems: totalItems,
	}

	return helper.ToGetUserResponsesWithPages(users), webPages
}

func (service *UserServiceImpl) UpdateById(ctx context.Context, request web.UserUpdateRequest, userId int, username string, isAdmin bool) web.UserUpdateOrCreateResponse {
	if !isAdmin {
		panic(exception.NewBadRequestErrors("you're not a super admin"))
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	service.UserRepository.Find(ctx, tx, userId, "")

	userRequest := domain.UserUpdateOrCreate{
		Id:       userId,
		Username: username,
		Name:     request.Name,
		Email:    request.Email,
	}

	userUpdate := service.UserRepository.Update(ctx, tx, userRequest, userId, "")
	return helper.ToUserUpdateOrCreateResponse(userUpdate)
}

func (service *UserServiceImpl) UpdateByUsername(ctx context.Context, request web.UserUpdateRequest, username string) web.UserUpdateOrCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	id := service.UserRepository.FindIdByUsername(ctx, tx, username)

	userRequest := domain.UserUpdateOrCreate{
		Id:       id,
		Username: username,
		Name:     request.Name,
		Email:    request.Email,
	}

	user := service.UserRepository.Update(ctx, tx, userRequest, 0, username)
	return helper.ToUserUpdateOrCreateResponse(user)
}

func (service *UserServiceImpl) DeleteById(ctx context.Context, userId int, isAdmin bool) {
	if !isAdmin {
		panic(exception.NewBadRequestErrors("you're not a super admin"))
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.UserRepository.Find(ctx, tx, userId, "")
	service.UserRepository.DeleteById(ctx, tx, userId)
}

func (service *UserServiceImpl) UpdatePasswordByUsername(ctx context.Context, request web.UserChangePasswordRequest, username string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	password := service.UserRepository.FindPassword(ctx, tx, username)
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(request.Password))
	if err != nil {
		panic(exception.NewBadRequestErrors("password wrong"))
	}

	if request.NewPassword != request.ReNewPassword {
		panic(exception.NewBadRequestErrors("new password doesn't match"))
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 10)
	helper.PanicIfError(err)
	userPassword := domain.User{Password: string(newPassword)}

	service.UserRepository.UpdatePassword(ctx, tx, userPassword, username, "")
}

func (service *UserServiceImpl) Logout(ctx context.Context, token string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tokenEncrypt := helper.EncryptData(token)

	blockedToken := domain.BlockedToken{Token: tokenEncrypt}

	service.UserRepository.Logout(ctx, tx, blockedToken)
}

func (service *UserServiceImpl) Verification(ctx context.Context, request web.UserVerificationRequest) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	digits := 8

	rand.Seed(time.Now().UnixNano())

	min := int64(1)
	max := int64(1)
	for i := 0; i < digits; i++ {
		max *= 10
	}

	randomNumber := rand.Int63n(max-min) + min

	name := service.UserRepository.FindByEmail(ctx, tx, request.Email)

	otpStr := strconv.FormatInt(randomNumber, 10)

	otps = append(otps, otpStr)
	encryptedOTP := helper.EncryptData(otpStr)

	otp := domain.OTP{
		Email: request.Email,
		OTP:   encryptedOTP,
	}

	service.UserRepository.SaveOTP(ctx, tx, otp)
	helper.VerifikasiUserEmail(request.Email, name, otpStr)
	return request.Email
}

func (service *UserServiceImpl) NewPassword(ctx context.Context, email string, request web.UserNewPasswordRequest) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	for _, otp := range otps {
		if otp != request.OTP {
			return errors.New("otp not valid")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 10)
	userRequest := domain.User{Password: string(hashedPassword)}

	service.UserRepository.UpdatePassword(ctx, tx, userRequest, "", email)
	return nil
}
