package user

import (
	"net/http"
	"prediksi-kelulusan-mahasiswa-naive-bayes/config"
	"prediksi-kelulusan-mahasiswa-naive-bayes/middleware"
	"prediksi-kelulusan-mahasiswa-naive-bayes/models"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	user := models.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed Login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed Login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{user.ID, user.Email, user.Name, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login User",
		"user":    userResponse,
	})

}

func Register(c echo.Context) error {

	// binding data
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create User",
		"User":    user,
	})
}
