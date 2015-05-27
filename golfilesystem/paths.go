package golfilesystem

import "os"

/*
PathExists is a non error-throwing simple boolean proxy for existence of a filesystem level path.
*/
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
