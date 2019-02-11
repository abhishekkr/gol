package golconv

import "strconv"

func StringToUint64(n string, defaultValue uint64) uint64 {
	val, err := strconv.Atoi(n)
	if err == nil {
		return uint64(val)
	}
	return defaultValue
}
