package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Ping Controller
type Ping struct{}

// New create Ping instance.
func (Ping) New() *Ping {
	return &Ping{}
}

// GetPing ....
func (controller *Ping) GetPing(c echo.Context) error {
	return c.JSON(http.StatusOK, `{"res":"pong"}`)
}
