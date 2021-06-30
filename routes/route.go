package routes

import (
	"net/http"
	"prediksi-kelulusan-mahasiswa-naive-bayes/controller"
	"prediksi-kelulusan-mahasiswa-naive-bayes/controller/profile_lulusan"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	// route group
	eRoute := e.Group("/api/v1")
	eRoute.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "api run teruss")
	})
	eRoute.GET("/profile-lulusan", profile_lulusan.Index)
	eRoute.GET("/prediksi-lulus", controller.PrediksiLulus)

	return e

}
