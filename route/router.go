package route

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
	"github.com/tomoyamachi/chabo/api"
	"github.com/tomoyamachi/chabo/db"
	myMw "github.com/tomoyamachi/chabo/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug()

	// Set MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Recover())
	e.Use(echoMw.Gzip())

	// Set Custom MiddleWare
	e.Use(myMw.TxMiddleware(db.Init()))

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.Get("/line", api.Line)
		v1.Get("/facebook", api.Facebook)
	}
	return e
}
