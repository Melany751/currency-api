package services

import (
	"net/http"
)

func OK(data any) (int, any) {
	return http.StatusOK, data
}

func Accepted(data any) (int, any) {
	return http.StatusAccepted, data
}

func Created(data any) (int, any) {
	return http.StatusCreated, data
}

func Updated(data any) (int, any) {
	return http.StatusOK, data
}

func Deleted(data any) (int, any) {
	return http.StatusNoContent, data
}

func NoContent() (int, any) {
	return http.StatusNoContent, nil
}

func SystemError(err any) (int, any) {
	return http.StatusInternalServerError, err
}

func BadRequest(err any) (int, any) {
	return http.StatusBadRequest, err
}

func NotFound(err any) (int, any) {
	return http.StatusNotFound, err
}
