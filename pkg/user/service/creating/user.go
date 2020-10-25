package creating

import (
	"github.com/GeorgVartanov/my_site/pkg/user/storage"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	FirstName string `json:"firstName" db:"firstName"`
	LastName  string `json:"lastName" db:"lastName"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`

}

func (u *User) ToUserStorageStruct() (*storage.User, error) {
	var MoscowTime, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, err
	}
	return &storage.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Created:   time.Now().In(MoscowTime),
		Changed:   time.Now().In(MoscowTime),
		IsAdmin:   false,
	}, nil
}

func (u *User) HashPassword() error  {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}