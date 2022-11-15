package app

import (
	"github.com/awryme/runtask/operations"
	"github.com/awryme/runtask/pkg/errors"
)

func defaultCmd(ctx Ctx) error {
	if len(ctx.Args) != 1 {
		return errors.Error("command to set as default not provided")
	}
	return operations.SetDefault(ctx.Logger, ctx.Args[0])
}
