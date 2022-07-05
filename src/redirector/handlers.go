package redirector

import (
	"github.com/labstack/echo/v4"
	"mango/pkg"
	"net/http"
)

type AppHandler struct {
	Storage pkg.IStorage
}

func (app AppHandler)Redirect(ctx echo.Context) error {
	hash_or_alias := ctx.Param("hash")
	item, err:= app.Storage.Check(hash_or_alias)
	if err!=nil{
		return ctx.String(http.StatusNotFound, "url does not exists")
	}
	return ctx.Redirect(http.StatusSeeOther, item.OriginalUrl)
}
