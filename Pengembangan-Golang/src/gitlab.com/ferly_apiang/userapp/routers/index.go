package routers

import (
	"echo-master"
	"net/http"
)

//index func
func Index() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Halaman Ini Di Kosongkan")
	})
	return e
}
