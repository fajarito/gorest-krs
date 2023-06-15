package krs_kec

type KrsRequest struct {
	IdKelurahan                         int
	KodeKelurahan                       string
	NamaKelurahan                       string
	IdKecamatan                         int
	KodeKecamatan                       string
	NamaKecamatan                       string
	IdKabupaten                         int
	KodeKabupaten                       string
	NamaKabupaten                       string
	IdProvinsi                          int
	KodeProvinsi                        string
	NamaProvinsi                        string
	JumlahKeluarga                      int
	JumlahKeluargaSasaran               int
	Prioritas1                          int
	Prioritas2                          int
	Prioritas3                          int
	Prioritas4                          int
	PeringkatKesejahteraanDiatas4       int
	JumlahKeluargaBeresikoStunting      int
	JumlahPus                           int
	JumlahPusHamil                      int
	JumlahBalita                        int
	JumlahBaduta                        int
	JumlahBukanPesertaKbModern          int
	JumlahJambanTidakLayak              int
	JumlahAirTidakLayak                 int
	JumlahKeluargaTidakBeresikoStunting int
}
