package handler

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

var ErrBind = NewAPIError(http.StatusBadRequest, "error binding to model")
var ErrForbidden = NewAPIError(http.StatusForbidden, "request forbidden")
var ErrServerError = NewAPIError(http.StatusInternalServerError, "unexpected error occured")

type APIError struct {
	eh *echo.HTTPError
}

func NewAPIError(code int, msg ...string) *APIError {
	return &APIError{eh: echo.NewHTTPError(code, msg...)}
}

func (e *APIError) Log(err error, msg ...string) *echo.HTTPError {
	logrus.WithFields(logrus.Fields{
		"code":  e.eh.Code(),
		"error": err,
		"msgs":  msg,
	}).Warn(e.eh.Error())
	return e.eh
}
