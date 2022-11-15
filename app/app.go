package app

import (
	"github.com/alecthomas/kong"
	"github.com/awryme/runtask/pkg/log"
	"github.com/awryme/runtask/pkg/stdio"
)

const desc = `
Run tasks or manipulate Taskfile. Action chosen by flag.

Default action: run specific task by name or run first task. Tasks can use special $cwd placeholder.`

type App struct {
	// general flags
	Debug bool `help:"enable debug logs"`
	Force bool `help:"on init command: force file creation"`

	// Args
	Args []string `arg:"" help:"command arguments" optional:""`

	// command flags
	Init    bool `help:"create taskfile in current dir" xor:"commands"`
	Set     bool `help:"set command: '--set <name> <full command>'" xor:"commands"`
	Default bool `help:"set command as default: '--default <name>'" xor:"commands"`
	Ls      bool `help:"list available commands" xor:"commands"`
	Rm      bool `help:"remove command: '--rm <name>'" xor:"commands"`
}

func Run() {
	var app App
	ctx := kong.Parse(&app, kong.UsageOnError(), kong.Name("runtask"), kong.Description(desc))
	command := chooseCommand(app)
	err := command(getCtx(app))
	ctx.FatalIfErrorf(err)
}

type Ctx struct {
	Debug   bool
	Force   bool
	Handles stdio.Handles
	Logger  log.Console
	Args    []string
}

func getCtx(app App) Ctx {
	handles := stdio.DefaultHandles()
	return Ctx{
		Debug:   app.Debug,
		Force:   app.Force,
		Handles: handles,
		Logger:  log.NewConsole(handles.Stderr, app.Debug),
		Args:    app.Args,
	}
}

type runFunc func(ctx Ctx) error

func chooseCommand(app App) runFunc {
	switch {
	case app.Init:
		return initCmd
	case app.Set:
		return setCmd
	case app.Ls:
		return lsCmd
	case app.Rm:
		return rmCmd
	case app.Default:
		return defaultCmd
	}
	return runCmd
}
