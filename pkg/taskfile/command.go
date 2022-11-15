package taskfile

import (
	"strings"
)

type Command struct {
	name string
	cmd  string
}

func (c Command) Name() string {
	return c.name
}

func (c Command) RawCommand() string {
	return c.cmd
}

func (c Command) Expand(cwd string) string {
	cmd := strings.ReplaceAll(c.cmd, "$cwd", cwd)
	return cmd
}
