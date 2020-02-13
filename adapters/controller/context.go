package controller

type Context interface {
	Param(name string) string
	Bind(i interface{}) error
	JSON(code int, i interface{}) error
}
