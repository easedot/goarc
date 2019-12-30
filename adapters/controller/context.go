package controller

import (
	"mime/multipart"
)

type Context interface {
	Param(name string) string
	QueryParam(name string) string
	FormValue(name string) string
	FormFile(name string) (*multipart.FileHeader, error)
	Bind(i interface{}) error
	JSON(code int,i interface{}) error
}