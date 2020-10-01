package storage

import (
	"fmt"
	"github.com/GeorgVartanov/my_site/pkg/user/adding"
	"time"
)

func (u *UserStorage) Create(user *adding.User) (int64, error) {
	var MoscowTime, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		return 0, err
	}

	newUser := User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Email:     user.Email,
		Created:   time.Now().In(MoscowTime),
		Changed:   time.Now().In(MoscowTime),
		IsAdmin:   false,
	}

	tx := u.DB.PostgresDB.MustBegin()

	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	res, err := tx.NamedExec(`INSERT INTO "myUser" ("firstName", "lastName", "email", "password", "created", "changed", "isAdmin") VALUES (:firstName, :lastName, :email, :password, :created, :changed, :isAdmin)`, &newUser)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0,err
	}
	id, err :=res.LastInsertId()
	if err != nil {
		return 0,err
	}
	fmt.Println(id)
	return id, nil
}
