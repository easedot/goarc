package router

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/easedot/hb_vendor/docs"
	"github.com/easedot/hb_vendor/domain"

	"github.com/easedot/hb_vendor/adapters/controller"
)

type ctxWrap struct {
	c echo.Context
}

func (f *ctxWrap) Get(name string) interface{} {
	return f.c.Get(name)
}
func (f *ctxWrap) Param(name string) string {
	return f.c.Param(name)
}
func (f *ctxWrap) FormValue(name string) string {
	return f.c.FormValue(name)
}
func (f *ctxWrap) FormFile(name string) (*multipart.FileHeader, error) {
	return f.c.FormFile(name)
}
func (f *ctxWrap) MultipartForm() (*multipart.Form, error) {
	return f.c.MultipartForm()
}

func (f *ctxWrap) Bind(i interface{}) error {
	return f.c.Bind(i)
}
func (f *ctxWrap) JSON(code int, i interface{}) error {
	return f.c.JSON(code, i)
}
func (f *ctxWrap) Stream(code int, contentType string, r io.Reader) error {
	return f.c.Stream(code, contentType, r)
}
func (f *ctxWrap) Attachment(file string, name string) error {
	return f.c.Attachment(file, name)
}
func (f *ctxWrap) File(file string) error {
	return f.c.File(file)
}

func urlSkipperProm(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/testurl") {
		return true
	}
	return false
}

func NewERouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	//e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"https://easedot.com", "https://easedot.net", "https://easedot.org"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	p := prometheus.NewPrometheus("echo", urlSkipperProm)
	p.Use(e)

	e.Static("/assets", "public")

	e.POST("/sign_in", func(context echo.Context) error { return c.SignIn(&ctxWrap{context}) })
	e.POST("/sign_up", func(context echo.Context) error { return c.SignUp(&ctxWrap{context}) })
	e.GET("/captcha/id", func(context echo.Context) error { return c.CaptchaGen(&ctxWrap{context}) })
	e.GET("/captcha/img/:id", func(context echo.Context) error { return c.CaptchaImg(&ctxWrap{context}) })

	v1 := e.Group("/api/v1")
	config := middleware.JWTConfig{
		Claims:     &domain.JwtCustomClaims{},
		SigningKey: []byte("hb_vendor_secret"),
	}
	v1.Use(middleware.JWTWithConfig(config))
	v1.GET("/vendors", func(context echo.Context) error { return c.GetVendors(&ctxWrap{c: context}) })
	v1.GET("/vendor/:id", func(context echo.Context) error { return c.GetVendor(&ctxWrap{c: context}) })
	v1.PUT("/vendor", func(context echo.Context) error { return c.UpdateVendor(&ctxWrap{c: context}) })
	v1.POST("/vendor", func(context echo.Context) error { return c.CreateVendor(&ctxWrap{c: context}) })

	v1.GET("/users", func(context echo.Context) error { return c.GetUsers(&ctxWrap{c: context}) })
	v1.POST("/user", func(context echo.Context) error { return c.CreateUser(&ctxWrap{c: context}) })
	v1.PUT("/user", func(context echo.Context) error { return c.UpdateUser(&ctxWrap{c: context}) })
	v1.PUT("/user/disable/:id", func(context echo.Context) error { return c.DisableUser(&ctxWrap{c: context}) })
	v1.PUT("/user/reset_password/:id", func(context echo.Context) error { return c.ResetPassword(&ctxWrap{c: context}) })

	v1.POST("/upload_file", func(context echo.Context) error { return c.UploadFiles(&ctxWrap{c: context}) })
	v1.GET("/upload_file/:id", func(context echo.Context) error { return c.GetUploadFiles(&ctxWrap{c: context}) })
	v1.DELETE("/upload_file/:id", func(context echo.Context) error { return c.DeleteUploadFiles(&ctxWrap{c: context}) })
	return e
}
