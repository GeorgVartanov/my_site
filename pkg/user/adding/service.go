package adding

// Repository ...
type Repository interface {
	Create(user *User) (int64, error)
}

type Service interface {
	Create(user *User) error
}


type service struct {
	Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(user *User) error{
	if _, err := s.Repository.Create(user); err!=nil{
		return  err
	}

	return nil
}