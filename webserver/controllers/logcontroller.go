package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yrisob/journal_log/config"
	"github.com/yrisob/journal_log/dto"
	"github.com/yrisob/journal_log/utils"
)

type LogController struct {
	Server IServer
}

// @Description получение логов системы
// @Param service query string true "один из вариантов: ppas, ppas_ehed, ppas_notifier"
// @Param since query string false "один из вариантов: week, yesterday, hour, minutes"
// @Tags logs
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Failure 405 {string} string
// @Failure 500 {string} string
// @Router /logs [get]
func (lc *LogController) GetLogs(c echo.Context) error {
	service := c.QueryParam("service")
	if !dto.IsNotServiceType(service) {
		return c.String(http.StatusBadRequest, "service должен быть одним из вариантов: ppas, ppas_ehed, ppas_notifier")
	}
	since := c.QueryParam("since")

	config, err := config.GetServerConfig()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log_service := config.GetPathByServiceType(service)
	sincePeriod := dto.GetSinceArgs(since)
	fmt.Printf("call cammand: $journalctl %s --since %s\n", log_service, sincePeriod)
	result, err := utils.CallSystemCommand("journalctl", config.GetPathByServiceType(service), "--since", dto.GetSinceArgs(since))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, result)

}

func (lc *LogController) AddRoutes() {
	lc.Server.GET("logs", lc.GetLogs)
}
