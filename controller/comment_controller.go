package controller

import (
	n "workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	v := new(n.Comments)

	if err := c.BodyParser(v); err != nil {
		return utility.ErrResponse(c, "Error in parsing", 400, err)
	}

	err := v.InsertComment()
	if err != nil {
		return utility.ErrResponse(c, "Error in Commenting", 400, err)
	}

	return c.JSON(fiber.Map{
		"message":  "Comment Successfully Added",
		"comments": v,
	})
}
