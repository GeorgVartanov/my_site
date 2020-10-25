package pgStorage

import (
	"github.com/GeorgVartanov/my_site/pkg/user/storage"
)

func (u *UserStorage) CreateUserRepository(user *storage.User) error {

	tx := u.DB.PostgresDB.MustBegin()

	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	_, err := tx.NamedExec(`INSERT INTO "myUser" ("firstName", "lastName", "email", "password", "created", "changed", "isAdmin") VALUES (:firstName, :lastName, :email, :password, :created, :changed, :isAdmin) RETURNING id`, &user)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	//fmt.Println(id)
	return nil
}
