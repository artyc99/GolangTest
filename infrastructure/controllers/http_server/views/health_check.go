package views

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MountHealthCheck(router *echo.Group) (err error) {
	router.GET("/health-check", func(eCtx echo.Context) error {
		return eCtx.JSON(http.StatusOK, struct{}{})
	})
	return nil
}
