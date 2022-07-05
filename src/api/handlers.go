package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	data "mango/data"
	"mango/pkg"
	"net/http"
	"strconv"
)


type AppHandler struct {
	Storage pkg.IStorage
}

// List godoc
// @Summary List available object.
// @Description Get list of uploaded urls.
// @Tags url
// @Success 200 {object} data.JSONResult{data=data.ResponseHashedData,code=int}
// @Failure 500 {object} data.JSONResult{data=interface{},code=int}
// @Router /api/url/ [get]
func (app AppHandler) List(ctx echo.Context) error {
	item, err := app.Storage.GetListItems()
	if err!=nil{
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, item)
}

// Retrieve godoc
// @Summary Retrieve url detail.
// @Description Get detail object by ID.
// @Param id path int64 true "Object id"
// @Tags url
// @Success 200 {object} data.JSONResult{data=data.ResponseHashedData,code=int}
// @Failure 400 {object} data.JSONResult{code=int,message=string}
// @Failure 404 {object} data.JSONResult{data=interface{},code=int,message=string}
// @Router /api/url/{id} [get]
func (app AppHandler) Retrieve(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		return ctx.JSON(http.StatusBadRequest, "url `id` is not provided")
	}
	item, err := app.Storage.GetItem(id)
	if err!=nil{
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, item)
}

// Put godoc
// @Summary Put url detail.
// @Description Change url data by ID.
// @Param id path int64 true "Object id"
// @Param url body data.HashedDataPutRequest true "Url Data"
// @Tags url
// @Success 200 {object} data.JSONResult{data=data.HashedData,code=int}
// @Failure 400 {object} data.JSONResult{code=int,data=interface{},message=string}
// @Failure 404 {object} data.JSONResult{data=interface{},code=int,message=string}
// @Router /api/url/{id} [put]
func (app AppHandler) Put(ctx echo.Context) error {
	var input data.HashedData

	id, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		return ctx.JSON(http.StatusBadRequest, "url `id` is not provided")
	}
	decoder := json.NewDecoder(ctx.Request().Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	item, err := app.Storage.UpdateItem(id, input)
	if err!=nil{
		return ctx.JSON(http.StatusOK, err)
	}

 	return ctx.JSON(http.StatusCreated, item)
}

// Create godoc
// @Summary Create url.
// @Description Create url shortner object from provided data.
// @Tags url
// @Accept json
// @Param url body data.HashedDataCreateRequest true "Url Data"
// @Success 201 {object} data.JSONResult{data=data.ResponseHashedData,code=int}
// @Failure 400 {object} data.JSONResult{code=int,data=interface{},message=string}
// @Failure 500 {object} data.JSONResult{data=interface{},code=int,message=string}
// @Router /api/url/ [post]
func (app AppHandler) Create(ctx echo.Context) error {
	var input data.HashedData

	decoder := json.NewDecoder(ctx.Request().Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	item, err := app.Storage.SetItem(input)
	if err!=nil{
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, item)
}

// Delete godoc
// @Summary Delete url.
// @Description Delete url shortner object by id.
// @Tags url
// @Param id path int64 true "Object id"
// @Success 200 {object} data.JSONResult{data=data.ResponseHashedData,code=int}
// @Failure 400 {object} data.JSONResult{code=int,data=interface{},message=string}
// @Failure 404 {object} data.JSONResult{data=interface{},code=int,message=string}
// @Router /api/url/ [delete]
func (app AppHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		return ctx.JSON(http.StatusBadRequest, "url `id` is not provided")
	}
	item, err := app.Storage.DeleteItem(id)
	if err!=nil{
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, item)
}