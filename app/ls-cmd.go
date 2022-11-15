package app

import (
	"github.com/awryme/runtask/operations"
)

func lsCmd(ctx Ctx) error {
	return operations.List(ctx.Logger, ctx.Handles)
}
