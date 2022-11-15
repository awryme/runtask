package operations

import (
	"github.com/awryme/runtask/constants"
	"github.com/awryme/runtask/pkg/log"
	"github.com/awryme/runtask/pkg/osfs"
)

func Remove(console log.Console, name string) error {
	path, err := osfs.FindUpward(constants.DefaultTaskFileName)
	if err != nil {
		return err
	}
	printUsingFile(console, path)

	fileData, err := readFile(path)
	if err != nil {
		return err
	}
	if !fileData.Delete(name) {
		console.Log("command %s doesn't exist", name)
		return nil
	}

	if err := writeFile(path, fileData); err != nil {
		return err
	}
	console.Log("removed command %s", name)
	return nil
}
