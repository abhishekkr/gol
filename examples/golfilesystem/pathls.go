package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
)

func prepareFS(basePath string) {
	childPath := fmt.Sprintf("%s/test-again", basePath)
	errDir := os.MkdirAll(childPath, 0755)
	if errDir != nil {
		panic(errDir)
	}

	d1 := []byte("hello\ngo\n")
	errFile := ioutil.WriteFile(fmt.Sprintf("%s/gol.log", basePath), d1, 0644)
	if errFile != nil {
		panic(errFile)
	}

	errFile = ioutil.WriteFile(fmt.Sprintf("%s/again.txt", childPath), d1, 0644)
	if errFile != nil {
		panic(errFile)
	}
}

func cleanFS(basePath string) {
	os.Remove(basePath)
}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	basePath := fmt.Sprintf("%s/gol-test-dir", dir)
	prepareFS(basePath)

	pathMap, err := golfilesystem.PathLs(basePath)
	if err != nil {
		panic(fmt.Sprintf("ERROR: PathWalk failed for '%s'", err.Error()))
	}
	if pathMap[basePath][0] != fmt.Sprintf("%s/gol.log", basePath) {
		panic(fmt.Sprintf("PathLs panic for gol.log!\n%q\n", pathMap))
	}
	if pathMap[basePath][1] != fmt.Sprintf("%s/test-again/again.txt", basePath) {
		panic(fmt.Sprintf("PathLs panic for again.txt!\n%q\n", pathMap))
	}

	pathMap, err = golfilesystem.PathLsN(basePath, 0)
	if err != nil {
		panic(fmt.Sprintf("ERROR: PathWalk failed for '%s'", err.Error()))
	}
	if pathMap[basePath][0] != fmt.Sprintf("%s/gol.log", basePath) {
		panic(fmt.Sprintf("PathLsN panic for gol.log!\n%q\n", pathMap))
	}
	if len(pathMap[basePath]) == 2 {
		panic(fmt.Sprintf("PathLsN panic for again.txt!\n%q\n", pathMap))
	}

	pathMap, err = golfilesystem.PathLsType(basePath, []string{".txt"})
	if err != nil {
		panic(fmt.Sprintf("ERROR: PathWalk failed for '%s'", err.Error()))
	}
	if len(pathMap[basePath]) == 2 {
		panic(fmt.Sprintf("PathLsType panic for gol.log!\n%q\n", pathMap))
	}
	if pathMap[basePath][0] != fmt.Sprintf("%s/test-again/again.txt", basePath) {
		panic(fmt.Sprintf("PathLsType panic for again.txt!\n%q\n", pathMap))
	}

	pathMap, err = golfilesystem.PathLsNType(basePath, 1, []string{".log"})
	if err != nil {
		panic(fmt.Sprintf("ERROR: PathWalk failed for '%s'", err.Error()))
	}
	if pathMap[basePath][0] != fmt.Sprintf("%s/gol.log", basePath) {
		panic(fmt.Sprintf("PathLsNType panic for gol.log!\n%q\n", pathMap))
	}
	if len(pathMap[basePath]) == 2 {
		panic(fmt.Sprintf("PathLsNType panic for again.txt!\n%q\n", pathMap))
	}

	pathMap, err = golfilesystem.PathLsTypeNot(basePath, []string{".txt"})
	if err != nil {
		panic(fmt.Sprintf("ERROR: PathWalk failed for '%s'", err.Error()))
	}
	if len(pathMap[basePath]) == 0 {
		panic(fmt.Sprintf("PathLsTypeNot panic for again.txt!\n%q\n", pathMap))
	}

	pathMap, err = golfilesystem.PathLsNTypeNot(basePath, 0, []string{".txt"})
	if err != nil {
		panic(fmt.Sprintf("ERROR: PathWalk failed for '%s'", err.Error()))
	}
	if len(pathMap[basePath]) == 0 {
		panic(fmt.Sprintf("PathLsNTypeNot panic for again.txt!\n%q\n", pathMap))
	}

	cleanFS(basePath)
}
