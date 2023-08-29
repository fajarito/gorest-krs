package krs_prov

// Buat Struct nya terlebih dahulu

type KrsProv struct {
	// gorm.Model

	IdKabupaten                         int64  `gorm:"column:id_kabupaten"`
	KodeKabupaten                       string `gorm:"column:kode_depdagri_kabupaten"`
	NamaKabupaten                       string `gorm:"column:nama_kabupaten"`
	IdProvinsi                          int64  `gorm:"column:id_provinsi"`
	KodeProvinsi                        string `gorm:"column:kode_depdagri_provinsi"`
	NamaProvinsi                        string `gorm:"column:nama_provinsi"`
	JumlahKeluarga                      int64  `gorm:"column:jumlah_keluarga"`
	JumlahKeluargaSasaran               int64  `gorm:"column:jumlah_keluarga_sasaran"`
	Prioritas1                          int64  `gorm:"column:prioritas_1"`
	Prioritas2                          int64  `gorm:"column:prioritas_2"`
	Prioritas3                          int64  `gorm:"column:prioritas_3"`
	Prioritas4                          int64  `gorm:"column:prioritas_4"`
	PeringkatKesejahteraanDiatas4       int64  `gorm:"column:bukan_prioritas"`
	JumlahKeluargaBeresikoStunting      int64  `gorm:"column:jml_klg_resiko_stunting"`
	JumlahPus                           int64  `gorm:"column:pus"`
	JumlahPusHamil                      int64  `gorm:"column:pus_hamil"`
	JumlahBalita                        int64  `gorm:"column:balita"`
	JumlahBaduta                        int64  `gorm:"column:baduta"`
	JumlahBukanPesertaKbModern          int64  `gorm:"column:bukan_peserta_kb_modern"`
	JumlahJambanTidakLayak              int64  `gorm:"column:jamban_layak_tidak"`
	JumlahAirTidakLayak                 int64  `gorm:"column:sumber_air_layak_tidak"`
	JumlahTerlaluBanyak                 int64  `gorm:"column:terlalu_banyak"`
	JumlahTerlaluMuda                   int64  `gorm:"column:terlalu_muda"`
	JumlahTerlaluTua                    int64  `gorm:"column:terlalu_tua"`
	JumlahTerlaluDekat                  int64  `gorm:"column:terlalu_dekat"`
	JumlahKeluargaTidakBeresikoStunting int64  `gorm:"column:tidak_risiko_stunting"`
}

func (KrsProv) TableName() string {
	return "sch_bkkbn.view_kabupaten_2022_krs"
}
