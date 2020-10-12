package creating

import "golang.org/x/crypto/bcrypt"

// UserCreatingRepo ...
type UserCreatingRepository interface {
	CreateUserRepository(user *User) error
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
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	if err := s.UserCreatingRepository.CreateUserRepository(user); err != nil {
		return err
	}
	return nil
}
