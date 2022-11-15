package app

import (
	"github.com/awryme/runtask/operations"
	"github.com/awryme/runtask/pkg/errors"
)

func runCmd(ctx Ctx) error {
	if len(ctx.Args) == 0 {
		return operations.Run(ctx.Logger, ctx.Handles, "")
	}
	if len(ctx.Args) > 1 {
		return errors.Error("commands don't take arguments")
	}
	return operations.Run(ctx.Logger, ctx.Handles, ctx.Args[0])
}
