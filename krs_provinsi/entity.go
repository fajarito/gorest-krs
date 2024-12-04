package krs_provinsi

// Buat Struct nya terlebih dahulu

type KrsProv struct {
	// gorm.Model

	Periode                             int64  `gorm:"column:periode"`
	IdProvinsi                          int64  `gorm:"column:id_provinsi"`
	KodeProvinsi                        string `gorm:"column:kode_provinsi"`
	NamaProvinsi                        string `gorm:"column:nama_provinsi"`
	JumlahKeluarga                      int64  `gorm:"column:jumlah_keluarga"`
	JumlahKeluargaSasaran               int64  `gorm:"column:jumlah_keluarga_sasaran"`
	Prioritas1                          int64  `gorm:"column:prioritas_1"`
	Prioritas2                          int64  `gorm:"column:prioritas_2"`
	Prioritas3                          int64  `gorm:"column:prioritas_3"`
	Prioritas4                          int64  `gorm:"column:prioritas_4"`
	PeringkatKesejahteraanDiatas4       int64  `gorm:"column:bukan_prioritas"`
	JumlahKeluargaBeresikoStunting      int64  `gorm:"column:total_prioritas"`
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
	TotalPus4T                          int64  `gorm:"column:total_pus_4t"`
	JumlahKeluargaTidakBeresikoStunting int64  `gorm:"column:tidak_risiko_stunting"`
}

func (KrsProv) TableName() string {
	return "public.agg_krs_provinsi"
}
