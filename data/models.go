// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package data

import (
	"database/sql"
)

type Todo struct {
	ID          int64
	Title       string
	Description sql.NullString
	CreatedAt   sql.NullTime
}
