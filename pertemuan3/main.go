package main

import (
	"pertemuan3/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello assalaam")
	})

	routes.StudentRoute(e)

	e.Logger.Fatal(e.Start(":8000"))
}