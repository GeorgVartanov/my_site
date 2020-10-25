package creating

import (
	"github.com/GeorgVartanov/my_site/pkg/user/storage"
)

// UserCreatingRepo ...
type UserCreatingRepository interface {
	CreateUserRepository(user *storage.User) error
}

type UserCreatingService interface {
	CreateUserService(user *User) error
}

type service struct {
	UserCreatingRepository
}

func NewService(r UserCreatingRepository) UserCreatingService {
	return &service{r}
}

func (s *service) CreateUserService(user *User) error {
	err:=user.HashPassword()
	if err != nil {
		return err
	}
	userStorage, err := user.ToUserStorageStruct()
	if err != nil {
		return err
	}
	if err := s.UserCreatingRepository.CreateUserRepository(userStorage); err != nil {
		return err
	}
	return nil
}
