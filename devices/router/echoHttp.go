package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/fasthttp/websocket"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang.org/x/time/rate"

	_ "github.com/easedot/goarc/docs"

	"github.com/easedot/goarc/adapters/controller"
)

type ctxWrap struct {
	c echo.Context
}

func (f *ctxWrap) Param(name string) string {
	return f.c.Param(name)
}
func (f *ctxWrap) Bind(i interface{}) error {
	return f.c.Bind(i)
}
func (f *ctxWrap) JSON(code int, i interface{}) error {
	return f.c.JSON(code, i)
}

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

func NewERouter(e *echo.Echo, c controller.AppController, rds *redis.Pool) *echo.Echo {
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

	e.Static("/assets", "public")

	v1 := e.Group("/api/v1")
	v1.GET("/articles", func(context echo.Context) error { return c.GetArticles(&ctxWrap{c: context}) })
	v1.GET("/article/:id", func(context echo.Context) error { return c.GetArticle(&ctxWrap{c: context}) })
	v1.PUT("/article/:id", func(context echo.Context) error { return c.UpdateArticle(&ctxWrap{c: context}) })
	v1.GET("/authors", func(context echo.Context) error { return c.GetAutuors(&ctxWrap{c: context}) })

	chat := NewChatServer(rds)
	e.GET("/ws", func(c echo.Context) error {
		var upgrader = websocket.Upgrader{}
		// Upgrade initial GET request to a websocket
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Printf("error: %v", err)
		}
		chat.ChatHandle(ws)
		return nil
	})

	return e
}

type RateLimitConfig struct {
	Skipper middleware.Skipper
	Path    []string
	Limit   int
	Burst   int
}

var Default = RateLimitConfig{
	Skipper: middleware.DefaultSkipper,
	Path:    []string{"*"},
	Limit:   2,
	Burst:   2,
}

func RateLimit() echo.MiddlewareFunc {
	return RateLimitWithConfig(Default)
}
func RateLimitWithConfig(r RateLimitConfig) echo.MiddlewareFunc {
	if r.Skipper == nil {
		r.Skipper = middleware.DefaultSkipper
	}
	if len(r.Path) == 0 {
		r.Path = Default.Path
	}

	lmt := rate.NewLimiter(rate.Limit(r.Limit), r.Burst)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if r.Skipper(c) {
				return next(c)
			}

			cfgPath := strings.Join(r.Path, "")
			reqPath := c.Path()
			if strings.Index(cfgPath, reqPath) > -1 {
				if !lmt.Allow() {
					return echo.ErrTooManyRequests
				}
			}
			return next(c)
		}
	}
}
