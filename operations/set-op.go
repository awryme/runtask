package operations

import (
	"github.com/awryme/runtask/constants"
	"github.com/awryme/runtask/pkg/log"
	"github.com/awryme/runtask/pkg/osfs"
)

func Set(console log.Console, name string, command string) error {
	path, err := osfs.FindUpward(constants.DefaultTaskFileName)
	if err != nil {
		return err
	}
	printUsingFile(console, path)

	fileData, err := readFile(path)
	if err != nil {
		return err
	}
	if cmd, ok := fileData.Get(name); ok {
		console.Log("command %s exists, previous command '%s'", name, cmd.RawCommand())
	}
	fileData.Set(name, command)

	if err := writeFile(path, fileData); err != nil {
		return err
	}
	console.Log("set command '%s' to '%s': OK", name, command)
	return nil
}
