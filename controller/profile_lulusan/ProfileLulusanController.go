package profile_lulusan

import (
	"net/http"
	"prediksi-kelulusan-mahasiswa-naive-bayes/config"
	"prediksi-kelulusan-mahasiswa-naive-bayes/models"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {

	var profileLulusan []models.ProfileLulusan

	err := config.DB.Find(&profileLulusan).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    profileLulusan,
	})

}
