package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// ReadRequest body and validate is
func ReadRequest(ctx *fiber.Ctx, request interface{}) error {
	if err := ctx.BodyParser(request); err != nil {
		return fmt.Errorf("error in ReadRequest in utils: %v", err.Error())
	}

	return validate.StructCtx(ctx.Context(), request)
}
