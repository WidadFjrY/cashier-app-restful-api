package helper

import (
	"github.com/widadfjry/cashier-app/model/domain"
	"github.com/widadfjry/cashier-app/model/web"
)

func ToUserResponse(user domain.User) web.UserResponses {
	return web.UserResponses{
		Id: user.Id,
		Data: web.GetDataUser{
			Username:  user.Username,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			AddedBy:   user.AddedBy,
			CratedAt:  user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}

func ToUserUpdateOrCreateResponse(user domain.UserUpdateOrCreate) web.UserUpdateOrCreateResponse {
	return web.UserUpdateOrCreateResponse{
		Id: user.Id,
		Data: web.UpdateDataUser{
			Username: user.Username,
			Name:     user.Name,
			Email:    user.Email,
		},
	}
}

func ToGetUserResponsesWithPages(users []domain.User) []web.UserResponses {
	var usersResponses []web.UserResponses
	for _, user := range users {
		usersResponses = append(usersResponses, ToUserResponse(user))
	}
	return usersResponses
}
