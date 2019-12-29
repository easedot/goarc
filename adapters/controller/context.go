package controller

type Context interface {
	Param(name string) string
	QueryParam(name string) string
	FormValue(name string) string
	Bind(i interface{}) error
	JSON(code int,i interface{}) error
}