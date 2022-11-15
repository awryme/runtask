package app

import "github.com/awryme/runtask/operations"

func initCmd(ctx Ctx) error {
	return operations.Init(ctx.Logger, ctx.Force)
}
