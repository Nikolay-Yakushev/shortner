package api

import (
	"github.com/labstack/echo/v4"
	"mango/pkg"
)


func RegisterRoutes(e *echo.Echo, storage pkg.IStorage) {
	appHandle := AppHandler{Storage: storage}
	group:= e.Group("/api")

	group.POST("/url/", appHandle.Create)
	group.GET("/url/", appHandle.List)
	group.GET("/url/:id", appHandle.Retrieve)
	group.DELETE("/url/:id", appHandle.Delete)
	group.PUT("/url/:id",  appHandle.Put)
}
