package osfs

import (
	"github.com/awryme/runtask/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func FindUpward(filename string) (ExistingFile, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return ExistingFile{}, errors.Wrap(err, "get cwd")
	}
	strPath, err := findFileCwd(cwd, filename)
	if err != nil {
		return ExistingFile{}, errors.Wrap(err, "find file")
	}
	path, err := NewPath(strPath)
	if err != nil {
		return ExistingFile{}, err
	}

	return NewExistingFile(path)
}

var osFileExists = func(file string) (bool, error) {
	info, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if info.IsDir() {
		return false, nil
	}
	return false, errors.Wrap(err, "stat file %s", file)
}

func findFileCwd(cwd string, filename string) (string, error) {
	dirs := pathParents(cwd)
	for _, dir := range dirs {
		path := filepath.Join(dir, filename)
		exists, err := osFileExists(path)
		if err != nil {
			return "", errors.Wrap(err, "file exists check")
		}
		if exists {
			return path, nil
		}
	}
	return "", errors.Wrap(os.ErrNotExist, "'%s' not found in dir '%s' or above", filename, cwd)
}

func pathParents(path string) []string {
	path = addSlash(path)
	parents := make([]string, 0)
	for path != "" {
		path = filepath.Dir(path)
		if path == "/" {
			path = ""
		}
		parents = append(parents, path)
	}
	return parents
}

func addSlash(s string) string {
	if strings.HasSuffix(s, "/") || strings.HasSuffix(s, "\\") {
		return s
	}
	return s + "/"
}
