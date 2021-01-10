package config

import (
	"fmt"
	"os"

	"github.com/yrisob/journal_log/dto"
)

//WebServerConfig - struct for storing web server config
type ServerConfig struct {
	Port         string
	User         string
	PPAS         string
	PPASEHED     string
	PPASNOTIFIER string
	PPASUPDATER  string
}

var serverConfig *ServerConfig

//GetServerConfig -  insert config from env params if they are exist or return error
func GetServerConfig() (*ServerConfig, error) {
	if serverConfig == nil {
		port, ok := os.LookupEnv("HTTP_PORT")
		if !ok {
			return nil, fmt.Errorf("HTTP_PORT is not defined")
		}

		user, ok := os.LookupEnv("OWNER")
		if !ok {
			return nil, fmt.Errorf("OWNER is not defined")
		}

		ppasPath, ok := os.LookupEnv("PPAS_PATH")
		if !ok {
			return nil, fmt.Errorf("PATH for ppas_dev is not defined")
		}

		ppasEhed, ok := os.LookupEnv("PPAS_EHED")
		if !ok {
			return nil, fmt.Errorf("PATH for ppas_ehed is not defined")
		}

		ppasNotifier, ok := os.LookupEnv("PPAS_NOTIFIER")
		if !ok {
			return nil, fmt.Errorf("PATH for ppas_notifier is not defined")
		}

		ppasUpdater, ok := os.LookupEnv("PPAS_UPDATER")
		if !ok {
			return nil, fmt.Errorf("PATH for ppas_updater is not defined")
		}

		serverConfig = &ServerConfig{
			Port:         port,
			User:         user,
			PPAS:         fmt.Sprintf(ppasPath, user),
			PPASEHED:     fmt.Sprintf(ppasEhed, user),
			PPASNOTIFIER: fmt.Sprintf(ppasNotifier, user),
			PPASUPDATER:  fmt.Sprintf(ppasUpdater, user),
		}
	}
	return serverConfig, nil

}

func (c *ServerConfig) GetPathByServiceType(serviceType dto.ServiceType) string {
	switch serviceType {
	case dto.PPAS_ehed:
		return c.PPASEHED
	case dto.PPAS_service:
		return c.PPAS
	case dto.PPAS_notifier:
		return c.PPASNOTIFIER
	case dto.PPAS_updater:
		return c.PPASUPDATER
	}
	return ""
}
