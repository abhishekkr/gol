package golfilesystem

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

/*
MkDir to make dir if not there already.
*/
func MkDir(dirpath string) error {
	if PathExists(dirpath) {
		return nil
	}

	err := os.MkdirAll(dirpath, 0755)
	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)
		return err
	}

	return nil
}

/*
CopyFile from src to dst
*/
func CopyFile(src, dst string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return nil
		}
	}
	if err = os.Link(src, dst); err == nil {
		return nil
	}

	if err = MkDir(path.Dir(dst)); err != nil {
		return err
	}
	err = copyFileContents(src, dst, sfi.Mode())
	return err
}

/*
copyFileContents copies the contents of the file named src to the file named
by dst. The file will be created if it does not already exist. If the
destination file exists, all it's contents will be replaced by the contents
of the source file.
*/
func copyFileContents(src string, dst string, srcMode os.FileMode) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	out.Chmod(srcMode)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
