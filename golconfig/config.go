package golconfig

/*FlatConfig can be used seamlessly between varied projects to share flatmap details around configuration.*/
type FlatConfig map[string]string

/*
Configurator enables provide config adapters for serializer.
*/
type Configurator interface {
	Unmarshal(configData string, config interface{})
	ConfigFromFile(configPath string, config interface{})
}

/*
ConfiguratorList contains reference to all config adapters to be fetched for usage.
*/
var ConfiguratorList = make(map[string]Configurator)

/*
RegisterConfigurator enables new (de)-serializing adapters to plug-in their reference.
*/
func RegisterConfigurator(name string, configType Configurator) {
	ConfiguratorList[name] = configType
}

/*
GetConfigurator lets client of this library fetch required adapter.
*/
func GetConfigurator(name string) Configurator {
	return ConfiguratorList[name]
}
