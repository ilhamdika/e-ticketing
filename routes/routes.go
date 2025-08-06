package routes

import (
	"e-ticketing/controllers"
	"e-ticketing/middleware"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.POST("/login", controllers.Login)

	protected := e.Group("")
	protected.Use(middleware.JWTMiddleware)
	protected.POST("/terminal", controllers.CreateTerminal)
}
