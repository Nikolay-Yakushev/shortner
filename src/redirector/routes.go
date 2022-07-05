package redirector

import (
	"github.com/labstack/echo/v4"
	"mango/pkg"
)

func RegisterRoutes(e *echo.Echo, storage pkg.IStorage) {
	appHandle := AppHandler{Storage: storage}
	e.GET("/:hash", appHandle.Redirect)

}

