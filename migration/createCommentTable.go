package migration

import (
	"context"
	"fmt"
)

// CreateRoleTable ...
func CreateCommentTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS comments (
		id serial PRIMARY KEY,
		comment VARCHAR ( 500 ) NOT NULL,
		user_id int,
		booking_id INTEGER REFERENCES bookings (id),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
	`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
