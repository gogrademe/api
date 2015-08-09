package handler

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

type APIError struct {
	*echo.HTTPError
}

func NewAPIError(code int, msg ...string) *APIError {
	return &APIError{echo.NewHTTPError(code, msg...)}
}

func (e *APIError) Log(err error, msg ...string) *echo.HTTPError {
	logrus.WithFields(logrus.Fields{
		"code":  e.Code(),
		"error": err,
	}).Warn(e.Error())
	return e.HTTPError
}
