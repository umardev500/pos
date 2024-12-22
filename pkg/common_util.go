package pkg

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

// BaseContext returns a context with a 5 second timeout
func BaseContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// ParseUnscopedParamContext parses the unscoped param and returns a new context
func ParseUnscopedParamContext(c *fiber.Ctx, ctx context.Context) context.Context {
	if !c.QueryBool("unscoped") {
		ctx = context.WithValue(ctx, UnscopedKey, false)
		return ctx
	}

	ctx = context.WithValue(ctx, UnscopedKey, true)
	return ctx
}

// ParseUnscopedParamContextBase is a shortcut for ParseUnscopedParamContext and BaseContext
func ParseUnscopedParamContextBase(c *fiber.Ctx) (context.Context, context.CancelFunc) {
	ctx, cancel := BaseContext()
	ctx = ParseUnscopedParamContext(c, ctx)
	return ctx, cancel
}
