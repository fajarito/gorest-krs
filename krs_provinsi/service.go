package krs_provinsi

type Service interface {
	KrsByNas(Periode int) ([]KrsProv, error)
	// KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, Periode int) ([]KrsKec, error)
	KrsProvById(KodeDepdagriProvinsi string, Periode int) ([]KrsProv, error)
	KrsAllByProv(periode int) ([]KrsProv, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) KrsByNas(Periode int) ([]KrsProv, error) {
	krses, err := s.repository.KrsByNas(Periode)
	return krses, err
}

func (s *service) KrsAllByProv(periode int) ([]KrsProv, error) {
	krses, err := s.repository.KrsAllByProv(periode)
	return krses, err
}

func (s *service) KrsProvById(KodeDepdagriProvinsi string, Periode int) ([]KrsProv, error) {
	krses, err := s.repository.KrsProvById(KodeDepdagriProvinsi, Periode)
	return krses, err
}
