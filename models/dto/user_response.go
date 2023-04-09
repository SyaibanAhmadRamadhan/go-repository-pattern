package dto

import "github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/models/entities"

type UsersResponse struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type UserTokenRespon struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func UserResponse(user entities.Users) UsersResponse {
	return UsersResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
