package controllers

import (
	"net/http"
	"time"

	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/helpers"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/models/dto"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/services"
	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserControllerImpl(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateReq := dto.UsersCreateRequest{}
	helpers.DecodeReq(r, &userCreateReq)

	res := controller.UserService.Create(r.Context(), userCreateReq)
	response := helpers.NewResponse(http.StatusCreated, "create user success", res)
	helpers.EncodeRes(w, response)
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userUpdateReq := dto.UsersUpdateRequest{}
	userUpdateReq.UpdatedAt = time.Now().Unix()
	helpers.DecodeReq(r, &userUpdateReq)

	userUpdateReq.Id = params.ByName("id")
	res := controller.UserService.Update(r.Context(), userUpdateReq)
	response := helpers.NewResponse(http.StatusOK, "update user success", res)
	helpers.EncodeRes(w, response)
}

func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	controller.UserService.Delete(r.Context(), userId)
	response := helpers.NewResponse(http.StatusOK, "delete user success", nil)
	helpers.EncodeRes(w, response)
}

func (controller *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	res := controller.UserService.FindById(r.Context(), userId)
	response := helpers.NewResponse(http.StatusOK, "data user", res)
	helpers.EncodeRes(w, response)
}

func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res := controller.UserService.FindAll(r.Context())
	response := helpers.NewResponse(http.StatusOK, "list data user", res)
	helpers.EncodeRes(w, response)
}
