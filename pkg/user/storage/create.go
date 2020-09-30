package storage

import (
	"github.com/GeorgVartanov/my_site/pkg/user/models"
	"log"
	"time"
)

func (u UserStorage) Create(user *models.User) error {
	var MoscowTime, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}

	user.Created = time.Now().In(MoscowTime)
	user.Changed = time.Now().In(MoscowTime)

	tx := u.DB.PostgresDB.MustBegin()

	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	_, err = tx.NamedExec(`INSERT INTO "myUser" ("firstName", "lastName", "email", "password", "created", "changed", "isAdmin") VALUES (:firstName, :lastName, :email, :password, :created, :changed, :isAdmin)`, &user)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
