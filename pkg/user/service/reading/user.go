package reading

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"firstName" db:"firstName"`
	LastName  string    `json:"lastName" db:"lastName"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Created   time.Time `json:"created" db:"created"`
	Changed   time.Time `json:"changed" db:"changed"`
	IsAdmin   bool      `json:"isAdmin" db:"isAdmin"`
}

type UserEmail struct {
	EmailName string `json:"email"`
}