package router

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

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
