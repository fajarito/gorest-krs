package krs_nas

type Service interface {
	KrsByNas() ([]KrsNas, error)
	KrsByProvDetail(KodeDepdagriProvinsi string) ([]KrsNas, error)
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

func (s *service) KrsByProvDetail(KodeDepdagriProvinsi string) ([]KrsNas, error) {
	faskeses, err := s.repository.KrsByProvDetail(KodeDepdagriProvinsi)
	return faskeses, err
}
