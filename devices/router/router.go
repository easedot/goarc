package router

import (
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/easedot/goarc/docs"

	"github.com/easedot/goarc/adapters/controller"
)

// @title Clean Arc API
// @version 1.0

// @contact.name API Support
// @contact.email easedot@gmail.com

// @host localhost:9090
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func urlSkipperProm(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/testurl") {
		return true
	}
	return false
}

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	//e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://easedot.com", "https://easedot.net", "https://easedot.org"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(RateLimitWithConfig(RateLimitConfig{
		Path:  []string{"/api/v1/articles"},
		Limit: 2,
		Burst: 4,
	}))

	p := prometheus.NewPrometheus("echo", urlSkipperProm)
	p.Use(e)

	//e.GET("/", func(context echo.Context) error {
	//	return context.String(http.StatusOK, "Hello, World!")
	//})
	e.GET("/doc/*", echoSwagger.WrapHandler)

	e.Static("/", "public")

	v1 := e.Group("/api/v1")
	v1.GET("/articles", func(context echo.Context) error { return c.GetArticles(context) })
	v1.GET("/article/:id", func(context echo.Context) error { return c.GetArticle(context) })
	v1.PUT("/article/:id", func(context echo.Context) error { return c.UpdateArticle(context) })
	v1.GET("/authors", func(context echo.Context) error { return c.GetAutuors(context) })
	return e
}
