package routes

import (
	"net/http"
	"prediksi-kelulusan-mahasiswa-naive-bayes/constants"
	"prediksi-kelulusan-mahasiswa-naive-bayes/controller"
	"prediksi-kelulusan-mahasiswa-naive-bayes/controller/profile_lulusan"
	"prediksi-kelulusan-mahasiswa-naive-bayes/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()

	// route group
	eRoute := e.Group("/api/v1")
	eRoute.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "api run teruss")
	})

	// middleware.LogMiddleware(e)
	eRoute.POST("/login", user.Login)
	eRoute.POST("/register", user.Register)
	eRoute.POST("/prediksi-lulus", controller.PrediksiLulus)

	eRoute.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eRoute.GET("/profile-lulusan", profile_lulusan.Index)
	eRoute.GET("/profile-lulusan/:nim", profile_lulusan.Show)
	eRoute.POST("/profile-lulusan", profile_lulusan.Store)
	eRoute.PUT("/profile-lulusan/:nim", profile_lulusan.Update)
	eRoute.DELETE("/profile-lulusan/:nim", profile_lulusan.Destroy)

	return e

}
