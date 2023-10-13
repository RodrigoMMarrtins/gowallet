package database

import (
	"database/sql"
	"net/mail"
)

type User struct {
	db *sql.DB
	ID     string
	Name   string
	WalletsID []string
	Email *mail.Address
	Password string
	TotalBalance float32
}
