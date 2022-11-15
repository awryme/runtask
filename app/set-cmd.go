package app

import (
	"github.com/awryme/runtask/operations"
	"github.com/awryme/runtask/pkg/errors"
)

func setCmd(ctx Ctx) error {
	if len(ctx.Args) != 2 {
		return errors.Error("command to set is not provided")
	}
	return operations.Set(ctx.Logger, ctx.Args[0], ctx.Args[1])
}
