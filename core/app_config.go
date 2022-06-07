package core

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// Singleton instance
var instance *AppConfig

type appConfigProps map[string]string

func readAppConfigProperty(filename string) (appConfigProps, error) {
	configProps := appConfigProps{}

	if len(filename) == 0 {
		return configProps, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				configProps[key] = value
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return configProps, nil
}

type applicationConfig struct {
	Name    string `json:"name" binding:"required"`
	Version string `json:"version"`
}

type serverConfig struct {
	Mode string `json:"mode"`
	Port string `json:"port"`
}

type AppConfig struct {
	Application applicationConfig `json:"application"`
	Server      serverConfig      `json:"server"`
}

/* GetAppConfig:
Returns the singleton instance of AppConfig
*/
func GetAppConfig() AppConfig {
	config := AppConfig{
		Application: applicationConfig{
			Name:    "Application",
			Version: "v1-SNAPSHOT",
		},
		Server: serverConfig{
			Mode: "DEVELOPMENT",
			Port: "8080",
		},
	}

	if nil == instance {

		configProps, err := readAppConfigProperty("application.properties")

		if val, found := configProps["application.name"]; found {
			config.Application.Name = val
		}

		if val, found := os.LookupEnv("APP_NAME"); found {
			config.Application.Name = val
		}

		if val, found := configProps["server.port"]; found {
			config.Server.Port = val
		}

		if val, found := os.LookupEnv("APP_PORT"); found {
			config.Server.Port = val
		}

		if val, found := configProps["server.mode"]; found {
			config.Server.Mode = val
		}

		if err != nil {
			log.Printf(err.Error())
			return config
		}
	}

	return config
}
