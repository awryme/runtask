package operations

import (
	"fmt"
	"github.com/awryme/runtask/constants"
	"github.com/awryme/runtask/pkg/errors"
	"github.com/awryme/runtask/pkg/log"
	"github.com/awryme/runtask/pkg/osfs"
	"os"
)

func Init(console log.Console, force bool) error {
	if force {
		return createFile(console)
	}
	path, err := osfs.FindUpward(constants.DefaultTaskFileName)
	if errors.Is(err, os.ErrNotExist) {
		return createFile(console)
	}
	if err != nil {
		return fmt.Errorf("try find existing file: %w", err)
	}
	return fmt.Errorf("file '%s' already exists", path.Rel())
}

func createFile(console log.Console) error {
	file, err := os.Create(constants.DefaultTaskFileName)
	if err != nil {
		return errors.Wrap(err, "create file")
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}
	console.Log("created file %s", file.Name())
	return nil
}
