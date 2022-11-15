package operations

import (
	"fmt"
	"github.com/awryme/runtask/constants"
	"github.com/awryme/runtask/pkg/errors"
	"github.com/awryme/runtask/pkg/log"
	"github.com/awryme/runtask/pkg/osfs"
	"github.com/awryme/runtask/pkg/runscript"
	"github.com/awryme/runtask/pkg/stdio"
	"os"
	"path/filepath"
)

func Run(console log.Console, handles stdio.Handles, name string) error {
	path, err := osfs.FindUpward(constants.DefaultTaskFileName)
	if err != nil {
		return err
	}
	printUsingFile(console, path)

	fileData, err := readFile(path)
	if err != nil {
		return errors.Wrap(err, "decode file")
	}

	cmd, ok := fileData.Get(name)
	if !ok {
		cmdLogName := "default command"
		if name != "" {
			cmdLogName = fmt.Sprintf("command '%s'", name)
		}
		return fmt.Errorf("%s not found", cmdLogName)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get cwd: %w", err)
	}
	expandedCommand := cmd.Expand(cwd)
	console.Debug("Full command: '%s'", expandedCommand)
	console.Debug("---")

	return execScript(handles, path, expandedCommand)
}

func execScript(handles stdio.Handles, taskfilePath osfs.ExistingFile, command string) error {
	workDir := filepath.Dir(taskfilePath.Full())

	return runscript.Run(handles, workDir, command)
}
