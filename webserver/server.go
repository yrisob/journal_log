package webserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/yrisob/journal_log/config"
	_ "github.com/yrisob/journal_log/docs"
	"github.com/yrisob/journal_log/webserver/controllers"
)

type IRouter interface {
	AddRoutes()
}

type (
	HttpServer struct {
		*echo.Echo
		controllers []IRouter
	}
)

// Validate - check data to struct properties conditions
func (s *HttpServer) Validate(data interface{}) error {
	return s.Validator.Validate(data)
}

// BindFormContext -  bind body param into input struct
func (s *HttpServer) BindFormContext(input interface{}, c echo.Context) error {
	if err := c.Bind(input); err != nil {
		return err
	}
	if err := s.Validate(input); err != nil {
		return err
	}
	return nil
}

//GetIntParamByName - get int param by name
func (s *HttpServer) GetIntParamByName(c echo.Context, name string) (uint, error) {
	sUID := c.Param(name)

	id, err := strconv.Atoi(sUID)

	return uint(id), err
}

//GetIDParam -  get uint id from  path param
func (s *HttpServer) GetIDParam(c echo.Context) (uint, error) {
	return s.GetIntParamByName(c, "id")
}

func (s *HttpServer) AddRoutes() {
	for _, controller := range s.controllers {
		controller.AddRoutes()
	}
}

// @title API для просмотра логов по девелоперским системам
// @version 1.0
// @description Набор методов для получения логов
// @termsOfService http://swagger.io/terms/
// @contact.name Соболев Юрий Андреевич
// @contact.url http://localhost:8100
// @license.name Apache 2.0
// @contact.email yrisob@gmail.com
func Start() error {
	config, err := config.GetServerConfig()
	if err != nil {
		return err
	}
	httpServer := &HttpServer{}
	httpServer.Echo = echo.New()
	httpServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodDelete}}))

	httpServer.controllers = []IRouter{
		&controllers.LogController{Server: httpServer},
	}
	httpServer.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
	httpServer.AddRoutes()

	httpServer.Logger.Fatal(httpServer.Start(fmt.Sprintf(":%s", config.Port)))
	return nil
}
