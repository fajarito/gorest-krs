package krk_kab

type Service interface {
	KrkByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, page int, pageSize int) ([]Krk, error)
	KrkByKecDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krk, error)
	GetTotalKrkCount(city string, district string) (int64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) KrkByKab(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, page int, pageSize int) ([]Krk, error) {

	krkes, err := s.repository.KrkByKab(KodeDepdagriProvinsi, KodeDepdagriKabupaten, page, pageSize)

	return krkes, err
}

func (s *service) KrkByKecDetail(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krk, error) {

	krkes, err := s.repository.KrkByKecDetail(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan)

	return krkes, err
}

func (s *service) GetTotalKrkCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string) (int64, error) {
	// Implement the logic to retrieve the total count of Krk items by city and district
	// Use the krkRepository to fetch the count from the database
	krktotal, err := s.repository.GetTotalKrkCount(KodeDepdagriProvinsi, KodeDepdagriKabupaten)
	return krktotal, err

}
