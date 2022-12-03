package stunting

type StuntingRequest struct {
	Nik                 string `binding:"required, number, len=16"`
	KodeProvinsi        int64
	NamaProvinsi        string
	KodeKabupaten       int64
	NamaKabupaten       string
	KodeKecamatan       int64
	NamaKecamatan       string
	KodeKelurahan       int64
	NamaKelurahan       string
	KodeRw              string
	KodeRt              string
	NamaKepalaKeluarga  string
	Baduta              string
	Balita              string
	Pus                 string
	PusHamil            string
	SumberAirLayakTidak string
	JambanLayakTidak    string
	TerlaluMuda         string
	TerlaluTua          string
	TerlaluDekat        string
	TerlaluBanyak       string
	RisikoStunting      string
}
