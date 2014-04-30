package golconfig

import (
	"encoding/json"
	"io/ioutil"
)

// init: register CSVMap to DataMap
func init() {
	RegisterConfig("json", new(ConfigJSON))
}

// send "json" to GetConfig to get it
type ConfigJSON struct{}

// populates the passed &map with config values from given file
func (config_json ConfigJSON) ConfigFromFile(config_path string, config interface{}) {
	config_bytes, err := ioutil.ReadFile(config_path)
	if err == nil {
		json.Unmarshal(config_bytes, config)
	}
}

// populates the passed &map with config values from given string
func (config_json ConfigJSON) Config(config_data string, config interface{}) {
	config_bytes := []byte(config_data)
	json.Unmarshal(config_bytes, config)
}
