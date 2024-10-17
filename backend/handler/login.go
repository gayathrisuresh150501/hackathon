package handler

import (
	"fmt"

	"github.com/savsgio/atreugo"
)

func LoginByEmailPassword(ctx *atreugo.RequestCtx) func(ctx *atreugo.RequestCtx) error {
	return func(ctx *atreugo.RequestCtx) error {
		name := string(ctx.PostArgs().Peek("name"))
		message := fmt.Sprintf("Hello, %s!", name)

		return ctx.TextResponse(message)
	}
}
