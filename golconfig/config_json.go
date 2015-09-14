package golconfig

import (
	"encoding/json"
	"io/ioutil"
)

/*
init registers ConfigJSON to Configurator.
*/
func init() {
	RegisterConfigurator("json", new(ConfigJSON))
}

/*
ConfigJSON send "json" to GetConfig to get it.
*/
type ConfigJSON struct{}

/*
ConfigFromFile populates the passed &map with config values from given file.
*/
func (configJSON ConfigJSON) ConfigFromFile(configPath string, config interface{}) {
	configBytes, err := ioutil.ReadFile(configPath)
	if err == nil {
		json.Unmarshal(configBytes, config)
	}
}

/*
Unmarshal populates the passed &map with config values from given string.
*/
func (configJSON ConfigJSON) Unmarshal(configData string, config interface{}) {
	configBytes := []byte(configData)
	json.Unmarshal(configBytes, config)
}
