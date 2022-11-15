package operations

import (
	"github.com/awryme/runtask/constants"
	"github.com/awryme/runtask/pkg/log"
	"github.com/awryme/runtask/pkg/osfs"
)

func SetDefault(console log.Console, name string) error {
	path, err := osfs.FindUpward(constants.DefaultTaskFileName)
	if err != nil {
		return err
	}
	printUsingFile(console, path)

	fileData, err := readFile(path)
	if err != nil {
		return err
	}
	if !fileData.SetDefault(name) {
		console.Log("command '%s' doesn't exist", name)
	}

	if err := writeFile(path, fileData); err != nil {
		return err
	}
	console.Log("set default command to '%s'", name)
	return nil
}
