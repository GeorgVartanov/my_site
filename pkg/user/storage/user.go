package storage

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

//CREATE TABLE "myUser" (
//"id" serial NOT NULL,
//"firstName" varchar(255) NOT NULL,
//"lastName" varchar(255) NOT NULL,
//"email" varchar(512) NOT NULL,
//"password" varchar(1024) NOT NULL,
//"created" TIMESTAMP NOT NULL,
//"changed" TIMESTAMP NOT NULL,
//"isAdmin" bool NOT NULL,
//CONSTRAINT "myUser_pk" PRIMARY KEY ("id")
//) WITH (
//OIDS=FALSE
//);