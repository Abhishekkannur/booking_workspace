package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

type CancelBooking struct {
	Id         int16     `json:"id"`
	CityId     int16     `json:"city_id"`
	BuildingId int16     `json:"building_id"`
	FloorId    int16     `json:"floor_id"`
	Purpose    string    `json:"purpose"`
	UserId     int16     `json:"user_id"`
	UserIds    []int16   `json:"user_ids"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (d *CancelBooking) DeleteBooking() error {

	//dt := time.Now()
	_, err := migration.DbPool.Exec(context.Background(), "DELETE FROM bookings WHERE id=$1", d.Id)
	fmt.Println(d.Id)

	return err

}
