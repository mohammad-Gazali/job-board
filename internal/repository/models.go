// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        int64
	Email     string
	Username  string
	Password  string
	CreatedAt pgtype.Timestamptz
}
