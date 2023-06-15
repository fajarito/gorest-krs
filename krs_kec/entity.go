package krs_kec

// Buat Struct nya terlebih dahulu

type Krs struct {
	// gorm.Model

	IdKelurahan                         int64  `gorm:"column:id_kelurahan"`
	KodeKelurahan                       string `gorm:"column:kode_depdagri_kelurahan"`
	NamaKelurahan                       string `gorm:"column:nama_kelurahan"`
	IdKecamatan                         int64  `gorm:"column:id_kecamatan"`
	KodeKecamatan                       string `gorm:"column:kode_depdagri_kecamatan"`
	NamaKecamatan                       string `gorm:"column:nama_kecamatan"`
	IdKabupaten                         int64  `gorm:"column:id_kabupaten"`
	KodeKabupaten                       string `gorm:"column:kode_depdagri_kabupaten"`
	NamaKabupaten                       string `gorm:"column:nama_kabupaten"`
	IdProvinsi                          int64  `gorm:"column:id_provinsi"`
	KodeProvinsi                        string `gorm:"column:kode_depdagri_provinsi"`
	NamaProvinsi                        string `gorm:"column:nama_provinsi"`
	JumlahKeluarga                      int64  `gorm:"column:jumlah_keluarga"`
	JumlahKeluargaSasaran               int64  `gorm:"column:jumlah_keluarga_sasaran"`
	Prioritas1                          int64  `gorm:"column:peringkat_kesejahteraan_1"`
	Prioritas2                          int64  `gorm:"column:peringkat_kesejahteraan_2"`
	Prioritas3                          int64  `gorm:"column:peringkat_kesejahteraan_3"`
	Prioritas4                          int64  `gorm:"column:peringkat_kesejahteraan_4"`
	PeringkatKesejahteraanDiatas4       int64  `gorm:"column:peringkat_kesejahteraan_diatas_4"`
	JumlahKeluargaBeresikoStunting      int64  `gorm:"column:jumlah_keluarga_beresiko_stunting"`
	JumlahPus                           int64  `gorm:"column:jumlah_pus"`
	JumlahPusHamil                      int64  `gorm:"column:jumlah_pus_hamil"`
	JumlahBalita                        int64  `gorm:"column:jumlah_balita"`
	JumlahBaduta                        int64  `gorm:"column:jumlah_baduta"`
	JumlahBukanPesertaKbModern          int64  `gorm:"column:jumlah_bukan_peserta_kb_modern"`
	JumlahJambanTidakLayak              int64  `gorm:"column:jumlah_jamban_tidak_layak"`
	JumlahAirTidakLayak                 int64  `gorm:"column:jumlah_air_layak_tidak"`
	JumlahKeluargaTidakBeresikoStunting int64  `gorm:"column:jumlah_keluarga_tidak_beresiko_stunting"`
}

func (Krs) TableName() string {
	return "sch_bkkbn.view_kelurahan_2022_krs"
}
