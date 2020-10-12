package storage

import (
	"github.com/GeorgVartanov/my_site/pkg/user/service/creating"
	"time"
)

func (u *UserStorage) CreateUserRepository(user *creating.User)  error {
	var MoscowTime, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		return  err
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
	_, err = tx.NamedExec(`INSERT INTO "myUser" ("firstName", "lastName", "email", "password", "created", "changed", "isAdmin") VALUES (:firstName, :lastName, :email, :password, :created, :changed, :isAdmin) RETURNING id`, &newUser)
	if err != nil {
		return  err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	//fmt.Println(id)
	return  nil
}
