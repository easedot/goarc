package controller

import (
	"io"
	"mime/multipart"
)

type Context interface {
	Param(name string) string
	Bind(i interface{}) error
	JSON(code int, i interface{}) error
	FormValue(name string) string
	FormFile(name string) (*multipart.FileHeader, error)
	MultipartForm() (*multipart.Form, error)
	Get(name string) interface{}
	Stream(code int, contentType string, r io.Reader) error
	Attachment(file string, name string) error
	File(file string) error
}
