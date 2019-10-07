package models

import "time"

//Todo db struct
type Todo struct {
	ID          int       `db:"id"`
	Description string    `db:"description"`
	DateCreated time.Time `db:"datecreated"`
}
