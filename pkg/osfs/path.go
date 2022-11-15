package osfs

import (
	"github.com/awryme/runtask/pkg/errors"
	"os"
	"path/filepath"
)

type Path interface {
	Rel() string
	Full() string
}

type osPath struct {
	relative string
	full     string
}

func (p osPath) Rel() string {
	return p.relative
}

func (p osPath) Full() string {
	return p.full
}

func NewPath(path string) (Path, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.Wrap(err, "get abs path")
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "get cwd")
	}
	rel, err := filepath.Rel(wd, abs)
	if err != nil {
		return nil, errors.Wrap(err, "get relative path")
	}

	return osPath{rel, abs}, nil
}
