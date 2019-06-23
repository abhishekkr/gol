package golconv

import "strconv"

func StringToInt(n string, defaultValue int) int {
	val, err := strconv.Atoi(n)
	if err == nil {
		return val
	}
	return defaultValue
}

func StringToUint64(n string, defaultValue uint64) uint64 {
	val, err := strconv.Atoi(n)
	if err == nil {
		return uint64(val)
	}
	return defaultValue
}

func StringToBool(b string, defaultValue bool) bool {
	switch b {
	case "true":
		return true
	case "false":
		return false
	}
	return defaultValue
}
