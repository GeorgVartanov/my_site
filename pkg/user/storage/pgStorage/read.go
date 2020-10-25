package pgStorage

import "github.com/GeorgVartanov/my_site/pkg/user/service/reading"

//CreateUserRepository
func (u *UserStorage) ReadUserRepository(email string) (*reading.User, error) {
	user := reading.User{}
	if err := u.DB.PostgresDB.Get(&user, `SELECT "id", "firstName", "lastName", "email", "created", "changed", "isAdmin" FROM "myUser" where "email"=$1` , email); err != nil {
		return nil, err
	}
	return &user, nil

}

func (u *UserStorage) ReadAllUserRepository() ([]reading.User, error) {
	var user []reading.User
	if err := u.DB.PostgresDB.Select(&user, `SELECT "id", "firstName", "lastName", "email", "created", "changed", "isAdmin" FROM "myUser"`); err != nil {
		return nil, err
	}
	return user, nil
}
