package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gopkg.in/ini.v1"
	"mango/data"
	_ "mango/docs/api"
	"mango/pkg"
	"mango/src/api"
	"mango/src/redirector"
	"net/http"
	"os"
)
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Success 200 {object} data.JSONResult{}
// @Router /hello [get]
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "It's working!")
}

func appRoutes(e *echo.Echo, storage pkg.IStorage) {
	e.GET("/hello", hello)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	api.RegisterRoutes(e, storage)
	redirector.RegisterRoutes(e, storage)
}

func getConfig(config_path string) data.Config {
	cfg, err := ini.Load(config_path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	conf := data.Config{}
	err = cfg.MapTo(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}

// @title Echo Swagger Example API
// @version 2.0
// @description This is Shortner service.
// @BasePath /
// @schemes http
func main() {
	e := echo.New()
	e.Use(mw.LoggerWithConfig(mw.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n"}))
	parser := argparse.NewParser("Shortner", "Parser for arguments")
	path := parser.String("p", "path",
		&argparse.Options{Required: false, Help: "Specify path to config.ini", Default: "./configs/config.ini"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	config:=getConfig(*path)
	dbErr:=pkg.CreateDBConnection(&config); if err!=nil{
		panic(dbErr)
	}
	storage, _ := pkg.NewDatabaseStorage()
	appRoutes(e, storage)
	addr:=fmt.Sprintf("%s:%s",config.Server.ServerHost, config.Server.ServerPort)
	e.Logger.Fatal(e.Start(addr))
}
