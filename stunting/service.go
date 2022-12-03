package stunting

type Service interface {
	FindByID(nik string) (Stunting, error)
	FindByPusHamil(nik string) (Stunting, error)
	FindByStunting(nik string) (Stunting, error)
	FindByBaduta(nik string) (Stunting, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByID(nik string) (Stunting, error) {
	stunting, err := s.repository.FindByID(nik)
	return stunting, err
}

func (s *service) FindByStunting(nik string) (Stunting, error) {
	stunting, err := s.repository.FindByStunting(nik)
	return stunting, err
}

func (s *service) FindByPusHamil(nik string) (Stunting, error) {
	stunting, err := s.repository.FindByPusHamil(nik)
	return stunting, err
}

func (s *service) FindByBaduta(nik string) (Stunting, error) {
	stunting, err := s.repository.FindByBaduta(nik)
	return stunting, err
}
