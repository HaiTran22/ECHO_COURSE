package handle

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func login(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.loginResponse{
		Token: "123456",
	})
}
