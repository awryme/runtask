package osfs

import (
	"fmt"
	"github.com/awryme/runtask/pkg/errors"
	"os"
)

type ExistingFile struct {
	path Path
}

func (e ExistingFile) Rel() string {
	return e.path.Rel()
}

func (e ExistingFile) Full() string {
	return e.path.Full()
}

func NewExistingFile(path Path) (ExistingFile, error) {
	info, err := os.Stat(path.Full())
	if err != nil {
		return ExistingFile{}, errors.Wrap(err, "stat file")
	}
	if info.IsDir() {
		return ExistingFile{}, fmt.Errorf("dir '%s' exists, should be file", path.Full())
	}
	return ExistingFile{path}, nil
}
