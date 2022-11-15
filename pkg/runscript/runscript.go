package runscript

import (
	"fmt"
	"github.com/awryme/runtask/pkg/stdio"
	"os"
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/parse"
	"src.elv.sh/pkg/shell"
)

func Run(handles stdio.Handles, dir string, cmd string) error {
	var fds = [3]*os.File{handles.Stdin, handles.Stdout, handles.Stderr}
	ev := shell.MakeEvaler(handles.Stderr)

	err := ev.Chdir(dir)
	if err != nil {
		return fmt.Errorf("eval: %w", err)
	}

	ports, cleanup := eval.PortsFromFiles(fds, ev.ValuePrefix())
	defer cleanup()

	src := parse.Source{
		Name:   "command",
		Code:   cmd,
		IsFile: false,
	}

	cfg := eval.EvalCfg{
		Ports:     ports,
		Interrupt: eval.ListenInterrupts,
		PutInFg:   true,
	}

	err = ev.Eval(src, cfg)
	if err != nil {
		return fmt.Errorf("eval: %w", err)
	}
	return nil
}
