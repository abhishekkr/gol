package golenv

import (
	"os"
	"strings"
)

func OverrideIfEnv(envVar string, defaultValue string) string {
	if HasEnv(envVar) {
		envVarValue := os.Getenv(envVar)
		if envVarValue != "" {
			return envVarValue
		}
	}
	return defaultValue
}

func HasEnv(envVar string) bool {
	for _, envKeyVal := range os.Environ() {
		if strings.Split(envKeyVal, "=")[0] == envVar {
			return true
		}
	}
	return false
}

func EnvMap() map[string]string {
	var key_val map[string]string
	key_val = make(map[string]string)

	for _, envKeyVal := range os.Environ() {
		envKeyValSplit := strings.Split(envKeyVal, "=")
		key_val[envKeyValSplit[0]] = envKeyValSplit[1]
	}
	return key_val
}
