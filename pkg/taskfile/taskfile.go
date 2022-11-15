package taskfile

type Taskfile struct {
	Commands []*Command
}

func (f *Taskfile) findCommand(name string) (*Command, int, bool) {
	for idx, command := range f.Commands {
		if command.name == name {
			return command, idx, true
		}
	}
	return nil, -1, false
}

func (f *Taskfile) Get(name string) (*Command, bool) {
	if name == "" {
		return f.Default()
	}
	cmd, _, ok := f.findCommand(name)
	return cmd, ok
}

func (f *Taskfile) Default() (*Command, bool) {
	if len(f.Commands) == 0 {
		return nil, false
	}
	return f.Commands[0], true
}

func (f *Taskfile) Set(name, cmd string) {
	command, _, ok := f.findCommand(name)
	if ok {
		command.cmd = cmd
		return
	}
	f.Commands = append(f.Commands, &Command{
		name: name,
		cmd:  cmd,
	})
}

func (f *Taskfile) Delete(name string) bool {
	_, idx, ok := f.findCommand(name)
	if !ok {
		return false
	}
	f.Commands = f.deletedCmd(idx)
	return true
}

func (f *Taskfile) SetDefault(name string) bool {
	cmd, idx, ok := f.findCommand(name)
	if !ok {
		return false
	}
	cmds := make([]*Command, 0, len(f.Commands))
	cmds = append(cmds, cmd)
	cmds = append(cmds, f.deletedCmd(idx)...)
	f.Commands = cmds
	return true
}

func (f *Taskfile) deletedCmd(idx int) []*Command {
	return append(f.Commands[:idx], f.Commands[idx+1:]...)
}
