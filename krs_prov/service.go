package krs_prov

type Service interface {
	KrsByProv(KodeDepdagriProvinsi string) ([]KrsProv, error)
	KrsByKabDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string) ([]KrsProv, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) KrsByProv(KodeDepdagriProvinsi string) ([]KrsProv, error) {
	faskeses, err := s.repository.KrsByProv(KodeDepdagriProvinsi)
	return faskeses, err
}

func (s *service) KrsByKabDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string) ([]KrsProv, error) {
	faskeses, err := s.repository.KrsByKabDetail(KodeDepdagriProvinsi, KodeDepdagriKabupaten)
	return faskeses, err
}
