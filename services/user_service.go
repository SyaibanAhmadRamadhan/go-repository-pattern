package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/helpers"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/models/dto"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/models/entities"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/repositories"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, request dto.UsersCreateRequest) dto.UsersResponse
	Update(ctx context.Context, request dto.UsersUpdateRequest) dto.UsersResponse
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) dto.UsersResponse
	FindAll(ctx context.Context) []dto.UsersResponse
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(user repositories.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: user,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, req dto.UsersCreateRequest) dto.UsersResponse {
	errValidate := service.Validate.Struct(req)
	helpers.PanicIfErr(errValidate)
	tx, err := service.DB.Begin()
	helpers.PanicIfErr(err)
	defer helpers.Defer(tx)
	uuid := uuid.New()
	passwordHash, _ := utils.HashPassword(req.Password)
	user := entities.Users{
		Id:        uuid.String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  passwordHash,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	user = service.UserRepository.Create(ctx, tx, user)
	return dto.UserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, req dto.UsersUpdateRequest) dto.UsersResponse {
	errValidate := service.Validate.Struct(req)
	helpers.PanicIfErr(errValidate)
	tx, err := service.DB.Begin()
	helpers.PanicIfErr(err)
	defer helpers.Defer(tx)

	user := service.UserRepository.FindById(ctx, tx, req.Id)

	user.Username = req.Username
	user.UpdatedAt = req.UpdatedAt

	user = service.UserRepository.Update(ctx, tx, user)

	return dto.UserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId string) {
	tx, err := service.DB.Begin()
	helpers.PanicIfErr(err)
	defer helpers.Defer(tx)

	service.UserRepository.Delete(ctx, tx, userId)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId string) dto.UsersResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfErr(err)
	defer helpers.Defer(tx)

	user := service.UserRepository.FindById(ctx, tx, userId)
	return dto.UserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []dto.UsersResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfErr(err)
	defer helpers.Defer(tx)

	var userResponses []dto.UsersResponse
	users := service.UserRepository.FindAll(ctx, tx)
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse(user))
	}
	return userResponses
}
