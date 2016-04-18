package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func Line(c echo.Context) error {
	//tx := c.Get("Tx").(*dbr.Tx)
	return c.String(http.StatusOK, "It is a Line access!\n")
}

func Facebook(c echo.Context) error {
	return c.String(http.StatusOK, "It is a facebook access!\n")
}
