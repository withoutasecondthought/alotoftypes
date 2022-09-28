package handler

import (
	"alotoftypes/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	employee := e.Group("/employee")
	{
		employee.GET("/:id", h.getEmployee)
		employee.POST("/", h.postEmployee)
		employee.PUT("/", h.putEmployee)
		employee.DELETE("/:id", h.deleteEmployee)
	}
	e.POST("/employees/:page", h.getEmployeesWithParams)

	return e
}
