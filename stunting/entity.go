package stunting

// Buat Struct nya terlebih dahulu

type Stunting struct {
	// gorm.Model

	Nik                 string `gorm:"column:nik"`
	KodeProvinsi        int64  `gorm:"column:kode_provinsi"`
	NamaProvinsi        string `gorm:"column:nama_provinsi"`
	KodeKabupaten       int64  `gorm:"column:kode_kabupaten"`
	NamaKabupaten       string `gorm:"column:nama_kabupaten"`
	KodeKecamatan       int64  `gorm:"column:kode_kecamatan"`
	NamaKecamatan       string `gorm:"column:nama_kecamatan"`
	KodeKelurahan       int64  `gorm:"column:kode_kelurahan"`
	NamaKelurahan       string `gorm:"column:nama_kelurahan"`
	KodeRw              string `gorm:"column:kode_rw"`
	KodeRt              string `gorm:"column:kode_rt"`
	Periode             int64  `gorm:"column:periode"`
	NamaKepalaKeluarga  string `gorm:"column:nama_kepala_keluarga"`
	Baduta              string `gorm:"column:baduta"`
	Balita              string `gorm:"column:balita"`
	Pus                 string `gorm:"column:pus"`
	PusHamil            string `gorm:"column:pus_hamil"`
	SumberAirLayakTidak string `gorm:"column:sumber_air_layak_tidak"`
	JambanLayakTidak    string `gorm:"column:jamban_layak_tidak"`
	TerlaluMuda         string `gorm:"column:terlalu_muda"`
	TerlaluTua          string `gorm:"column:terlalu_tua"`
	TerlaluDekat        string `gorm:"column:terlalu_dekat"`
	TerlaluBanyak       string `gorm:"column:terlalu_banyak"`
	RisikoStunting      string `gorm:"column:risiko_stunting"`
}

func (Stunting) TableName() string {
	return "sch_bkkbn.mv_krs"
}
