package middlewares

import "github.com/gofiber/fiber/v2"

func GetClaim(c *fiber.Ctx) *Claim {
	claim, ok := c.Locals("claim").(*Claim)
	if !ok {
		return nil
	}
	return claim
}
