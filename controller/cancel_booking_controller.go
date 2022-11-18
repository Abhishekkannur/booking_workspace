package controller

import (
	"strconv"
	d "workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

func DeleteBookings(c *fiber.Ctx) error {
	id := c.Params("id")

	j, _ := strconv.Atoi(id)

	u := &d.CancelBooking{Id: int16(j)}

	err := u.DeleteBooking()

	if err != nil {
		return utility.ErrResponse(c, "Error in canceling bookings", 400, err)
	}

	return c.JSON(fiber.Map{
		"message": "Bookings Successfully Canceled",
	})
}
