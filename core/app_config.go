package core

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type applicationConfig struct {
	Name    string `json:"name" binding:"required"`
	Version string `json:"version"`
}

type serverConfig struct {
	Mode     string `json:"mode"`
	Port     string `json:"port"`
	BasePath string `json:"base_path"`
}

type AppConfig struct {
	Application applicationConfig `json:"application"`
	Server      serverConfig      `json:"server"`
}

/* GetAppConfig:
Returns the singleton instance of AppConfig
*/

var logger = GetLogger("core")

// Singleton instance
var instance *AppConfig

type appConfigProps map[string]string

func readAppConfigProperty(filename string) (appConfigProps, error) {
	logger.Tracef("entry: readAppConfigProperty(%s)", filename)
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

	logger.Tracef("exit: readAppConfigProperty()")
	return configProps, nil
}

func ExistOrDefault(value string, defaultValue string) string {
	trimmedValue := strings.Trim(value, " ")
	if len(trimmedValue) > 0 {
		return trimmedValue
	}

	return defaultValue
}

func ConfigExistOrElse(configMap map[string]string, key string, envKey string, defaultValue string) string {
	var value string = defaultValue
	if val, found := configMap[key]; found {
		logger.Tracef("Config[%s] found as '%s'\n", key, val)
		value = val
	} else if val, found := os.LookupEnv(envKey); found {
		logger.Tracef("ENV[%s] found as '%s'\n", envKey, val)
		value = val
	}

	return ExistOrDefault(value, defaultValue)
}

func GetAppConfig() AppConfig {

	if nil == instance {

		configProps, err := readAppConfigProperty("application.properties")
		if err != nil {
			logger.Warnf("Error while reading config from application.properties")
		}

		instance = &AppConfig{
			Application: applicationConfig{
				Name:    ConfigExistOrElse(configProps, "application.name", "APP_NAME", "Application"),
				Version: ConfigExistOrElse(configProps, "application.version", "APP_VERSION", "0.0.1-snapshot"),
			},
			Server: serverConfig{
				Mode:     ConfigExistOrElse(configProps, "server.mode", "SERVER_MODE", "DEVELOPMENT"),
				Port:     ConfigExistOrElse(configProps, "server.port", "SERVER_PORT", "8080"),
				BasePath: ConfigExistOrElse(configProps, "server.base_path", "SERVER_BASE_PATH", ""),
			},
		}

	}

	return *instance
}
