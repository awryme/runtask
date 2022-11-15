package operations

import (
	"fmt"
	"github.com/awryme/runtask/constants"
	"github.com/awryme/runtask/pkg/log"
	"github.com/awryme/runtask/pkg/osfs"
	"github.com/awryme/runtask/pkg/stdio"
	"github.com/awryme/runtask/pkg/taskfile"
)

func List(console log.Console, handles stdio.Handles) error {
	path, err := osfs.FindUpward(constants.DefaultTaskFileName)
	if err != nil {
		return err
	}
	printUsingFile(console, path)

	fileData, err := readFile(path)
	if err != nil {
		return err
	}
	for _, command := range fileData.Commands {
		printCommand(handles, command)
	}
	return nil
}

func printCommand(handles stdio.Handles, command *taskfile.Command) {
	_, _ = fmt.Fprintf(handles.Stdout, "%s: %s\n", command.Name(), command.RawCommand())
}
