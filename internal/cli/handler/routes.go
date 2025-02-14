package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) PublicRoutes(e *echo.Echo) {
	e.GET("/", h.MainPage)
	e.GET("/auth", h.GetAuthPage)
	e.POST("/auth", h.Authenticate)
	e.GET("/exit", h.Exit)
	e.GET("/act", h.Action)
	e.GET("/history", h.History)
	e.GET("/servers", h.Servers)
	e.HEAD("/", h.MainPage)
}
