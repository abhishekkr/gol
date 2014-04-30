package golconfig

type Config interface {
	Config(config_data string, config interface{})
	ConfigFromFile(config_path string, config interface{})
}

var ConfigList = make(map[string]Config)

func RegisterConfig(name string, configType Config) {
	ConfigList[name] = configType
}

func GetConfig(name string) Config {
	return ConfigList[name]
}
