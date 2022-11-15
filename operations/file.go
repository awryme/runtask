package operations

import (
	"github.com/awryme/runtask/constants"
	"github.com/awryme/runtask/pkg/errors"
	"github.com/awryme/runtask/pkg/log"
	"github.com/awryme/runtask/pkg/osfs"
	"github.com/awryme/runtask/pkg/taskfile"
	"os"
)

func readFile(path osfs.ExistingFile) (*taskfile.Taskfile, error) {
	file, err := os.Open(path.Full())
	if err != nil {
		return nil, errors.Wrap(err, "open file")
	}
	defer file.Close()

	fileData, err := taskfile.Decode(file)
	if err != nil {
		return nil, errors.Wrap(err, "decode file")
	}
	return fileData, nil
}

func writeFile(path osfs.ExistingFile, data *taskfile.Taskfile) error {
	file, err := os.OpenFile(path.Full(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "create file")
	}
	defer file.Close()

	err = taskfile.Encode(file, data)
	if err != nil {
		return errors.Wrap(err, "decode file")
	}
	return nil
}

func printUsingFile(console log.Console, path osfs.ExistingFile) {
	if constants.DefaultTaskFileName != path.Rel() {
		console.Log("Using file %s", path.Rel())
	} else {
		console.Debug("Using file %s", path.Full())
	}
}
