package storage

func (u *UserStorage) Read(email string) (User, error) {
	user := User{}
	if err := u.DB.PostgresDB.Get(&user, `SELECT "id", "firstName", "lastName", "email", "created", "changed", "isAdmin" FROM "myUser" where "email"=$1` , email); err != nil {
		return user, err
	}
	return user, nil

}

func (u *UserStorage) ReadAll() ([]User, error) {
	var user []User
	if err := u.DB.PostgresDB.Select(&user, `SELECT "id", "firstName", "lastName", "email", "created", "changed", "isAdmin" FROM "myUser"`); err != nil {
		return nil, err
	}
	return user, nil
}
