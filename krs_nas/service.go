package krs_nas

type Service interface {
	KrsByNas() ([]KrsNas, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) KrsByNas() ([]KrsNas, error) {
	krsnases, err := s.repository.KrsByNas()
	return krsnases, err
}
