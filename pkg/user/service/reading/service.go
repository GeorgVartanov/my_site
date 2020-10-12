package reading

// UserReadingRepository ...
type UserReadingRepository interface {
	ReadUserRepository(email string) (*User, error)
	ReadAllUserRepository() ([]User, error)
}

type UserReadingService interface {
	ReadUserService(email string) (*User, error)
}

type service struct {
	UserReadingRepository
}

func NewService(r UserReadingRepository) UserReadingService {
	return &service{r}
}

func (s *service) ReadUserService(email string) (*User, error) {
	user, err := s.UserReadingRepository.ReadUserRepository(email)
	if  err != nil {
		return nil, err
	}
	return user, nil
}



func (s *service) ReadallUserService(email string) ([]User, error) {
	user, err := s.UserReadingRepository.ReadAllUserRepository()
	if  err != nil {
		return nil, err
	}
	return user, nil
}