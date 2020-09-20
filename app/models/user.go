package models

import "time"

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	RoleID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
