package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Hello
// @Tags	example
// @id		Hello
// @Accept	json
// @Produce	json
// @Param	c query echo.Context true "Context"
// @Success	200
// @Router	/hello [get]
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Goodbye ,
// @Tags	example
// @id		Goodbye
// @Accept	json
// @Produce	json
// @success	200
// @Param	c query echo.Context true "Context"
// @Router	/goodbye [get]
func Goodbye(c echo.Context) error {
	return c.String(http.StatusOK, "Goodbye, World!")
}
