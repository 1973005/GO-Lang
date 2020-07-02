package main

import (
	"echo-master"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Halaman Ini Di Kosongkan")
	})
	e.Start(":3000")
}
