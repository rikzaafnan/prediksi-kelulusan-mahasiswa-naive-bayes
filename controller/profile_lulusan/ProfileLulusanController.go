package profile_lulusan

import (
	"net/http"
	"prediksi-kelulusan-mahasiswa-naive-bayes/config"
	"prediksi-kelulusan-mahasiswa-naive-bayes/models"
	"strconv"

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

func Show(c echo.Context) error {

	var profileLulusan models.ProfileLulusan

	profileLulusanNIM, _ := strconv.Atoi(c.Param("nim"))

	err := config.DB.First(&profileLulusan, profileLulusanNIM).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "Profile Alumni",
		"profile_lulusan": profileLulusan,
	})

}

func Store(c echo.Context) error {

	var profileLulusan models.ProfileLulusan
	c.Bind(&profileLulusan)

	err := config.DB.Create(&profileLulusan).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"text":    "Gagal",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "Success tambah data Alumni",
		"profile_lulusan": profileLulusan,
	})

}

func Update(c echo.Context) error {

	var profileLulusan models.ProfileLulusan

	profileLulusanNim, _ := strconv.Atoi(c.Param("nim"))

	c.Bind(&profileLulusan)

	err := config.DB.Where("nim = ?", profileLulusanNim).Updates(&profileLulusan).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"text":    "Gagal",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "Success update data Alumni",
		"profile_lulusan": profileLulusan,
	})

}

func Destroy(c echo.Context) error {

	var profileLulusan models.ProfileLulusan

	profileLulusanNim, _ := strconv.Atoi(c.Param("nim"))

	c.Bind(&profileLulusan)

	err := config.DB.Where("nim = ?", profileLulusanNim).Delete(&profileLulusan).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"text":    "Gagal",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success hapus data Alumni",
	})

}
