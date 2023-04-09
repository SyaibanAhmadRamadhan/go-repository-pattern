package exception

import (
	"net/http"

	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/helpers"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	} else if badRequestError(w, r, err) {
		return
	} else if validationErrors(w, r, err) {
		return
	} else {
		internalServerError(w, r, err)
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := helpers.NewResponse(http.StatusInternalServerError, "INTERNAL SERVER ERROR", err)
	helpers.EncodeRes(w, response)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response := helpers.NewResponse(http.StatusNotFound, "NOT FOUND", exception)
		helpers.EncodeRes(w, response)
		return true
	} else {
		return false
	}
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := helpers.NewResponse(http.StatusBadRequest, "BAD REQUEST", exception)
		helpers.EncodeRes(w, response)
		return true
	} else {
		return false
	}
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := helpers.NewResponse(http.StatusBadRequest, "BAD REQUEST", helpers.NewResponValidate(exception))
		helpers.EncodeRes(w, response)
		return true
	} else {
		return false
	}
}
