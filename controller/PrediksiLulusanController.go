package controller

import (
	"net/http"
	"prediksi-kelulusan-mahasiswa-naive-bayes/config"
	"prediksi-kelulusan-mahasiswa-naive-bayes/models"

	"github.com/labstack/echo/v4"
)

func totalDataStatusKelulusan(keteranganStatusKelulusan string) ([]models.ProfileLulusan, error) {

	var profileLulusan []models.ProfileLulusan

	err := config.DB.Where("status_kelulusan = ?", keteranganStatusKelulusan).Find(&profileLulusan).Error
	if err != nil {
		return profileLulusan, err
	}

	return profileLulusan, nil
}

func PrediksiLulus(c echo.Context) error {

	// peluang prediksi Status Mahasiswa ( TEPAT atau tidak TEPAT )
	var profileLulusan []models.ProfileLulusan

	// mencari total data training
	err := config.DB.Find(&profileLulusan).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	totalDataLulusan := len(profileLulusan)

	// // total data 'TEPAT' yang ada pada data training
	totalDataTepat, err := totalDataStatusKelulusan("TEPAT")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	validTotalDataLulusanTepat := float32(len(totalDataTepat)) / float32(totalDataLulusan)

	// // total data 'TERLAMBAT' yang ada pada data training
	totalDataTerlambat, err := totalDataStatusKelulusan("TERLAMBAT")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	validTotalDataLulusanTerlambat := float32(len(totalDataTerlambat)) / float32(totalDataLulusan)

	response := map[string]interface{}{
		"total-data-lulusan":        totalDataLulusan,
		"total-data-tepat":          len(totalDataTepat),
		"lulusan-tepat / total":     validTotalDataLulusanTepat,
		"total-data-terlambat":      len(totalDataTerlambat),
		"lulusan-terlambat / total": validTotalDataLulusanTerlambat,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"response": response,
	})

}
