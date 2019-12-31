package golenv

import (
	"os"
	"strings"

	"github.com/abhishekkr/gol/golconv"
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

func OverrideIfEnvBool(envVar string, defaultValue bool) bool {
	if HasEnv(envVar) {
		envVarValue := os.Getenv(envVar)
		return golconv.StringToBool(envVarValue, defaultValue)
	}
	return defaultValue
}

func OverrideIfEnvInt(envVar string, defaultValue int) int {
	if HasEnv(envVar) {
		envVarValue := os.Getenv(envVar)
		return golconv.StringToInt(envVarValue, defaultValue)
	}
	return defaultValue
}

func OverrideIfEnvUint64(envVar string, defaultValue uint64) uint64 {
	if HasEnv(envVar) {
		envVarValue := os.Getenv(envVar)
		return golconv.StringToUint64(envVarValue, defaultValue)
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
	var key_val = make(map[string]string)

	for _, envKeyVal := range os.Environ() {
		envKeyValSplit := strings.Split(envKeyVal, "=")
		key_val[envKeyValSplit[0]] = envKeyValSplit[1]
	}
	return key_val
}
