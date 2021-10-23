package product

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) FindProductByID(id int, userID int) (*Product, error) {
	return s.repo.FindProductByID(id, userID)
}
