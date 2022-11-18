package model

import (
	"context"
	"time"
	"workspace_booking/migration"
)

type Comments struct {
	Id        int       `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comments) InsertComment() error {
	dt := time.Now()
	query := "INSERT INTO comments (id , comment, " +
		"user_id, created_at, updated_at)" +
		" VALUES ($1, $2,$3,$4,$5) RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(
		context.Background(), query, c.Id, c.Comment, dt, dt,
	)
	err := d.Scan(&c.Id, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
