package krs_nas

type KrsNasRequest struct {
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
