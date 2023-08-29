package krs_nas

type KrsNasResponse struct {
	IdProvinsi                          int    `json:"id_provinsi"`
	KodeProvinsi                        string `json:"kode_provinsi"`
	NamaProvinsi                        string `json:"nama_provinsi"`
	JumlahKeluarga                      int    `json:"jumlah_keluarga"`
	JumlahKeluargaSasaran               int    `json:"jumlah_keluarga_sasaran"`
	Prioritas1                          int    `json:"peringkat_kesejahteraan_1"`
	Prioritas2                          int    `json:"peringkat_kesejahteraan_2"`
	Prioritas3                          int    `json:"peringkat_kesejahteraan_3"`
	Prioritas4                          int    `json:"peringkat_kesejahteraan_4"`
	PeringkatKesejahteraanDiatas4       int    `json:"peringkat_kesejahteraan_diatas_4"`
	JumlahKeluargaBeresikoStunting      int    `json:"jumlah_keluarga_beresiko_stunting"`
	JumlahPus                           int    `json:"jumlah_pus"`
	JumlahPusHamil                      int    `json:"jumlah_pus_hamil"`
	JumlahBalita                        int    `json:"jumlah_balita"`
	JumlahBaduta                        int    `json:"jumlah_baduta"`
	JumlahBukanPesertaKbModern          int    `json:"jumlah_bukan_peserta_kb_modern"`
	JumlahJambanTidakLayak              int    `json:"jumlah_jamban_tidak_layak"`
	JumlahAirTidakLayak                 int    `json:"jumlah_air_tidak_layak"`
	JumlahTerlaluBanyak                 int    `json:"jumlah_terlalu_banyak"`
	JumlahTerlaluMuda                   int    `json:"jumlah_terlalu_muda"`
	JumlahTerlaluTua                    int    `json:"jumlah_terlalu_tua"`
	JumlahTerlaluDekat                  int    `json:"jumlah_terlalu_dekat"`
	JumlahKeluargaTidakBeresikoStunting int    `json:"jumlah_keluarga_tidak_beresiko_stunting"`
}
