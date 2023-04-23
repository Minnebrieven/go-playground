package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "user_agent=${user_agent} method=${method}, uri=${uri}, status=${status} latency_human=${latency_human}\n",
	}))

}
