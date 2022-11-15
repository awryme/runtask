package app

import (
	"github.com/awryme/runtask/operations"
	"github.com/awryme/runtask/pkg/errors"
)

func rmCmd(ctx Ctx) error {
	if len(ctx.Args) != 1 {
		return errors.Error("command to remove is not provided")
	}
	return operations.Remove(ctx.Logger, ctx.Args[0])
}
