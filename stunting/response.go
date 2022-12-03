package stunting

type StuntingResponse struct {
	Nik                 string `json:"nik"`
	KodeProvinsi        int64  `json:"kode provinsi"`
	NamaProvinsi        string `json:"nama provinsi"`
	KodeKabupaten       int64  `json:"kode kabupaten"`
	NamaKabupaten       string `json:"nama kabupaten"`
	KodeKecamatan       int64  `json:"kode kecamatan"`
	NamaKecamatan       string `json:"nama kecamatan"`
	KodeKelurahan       int64  `json:"kode kelurahan"`
	NamaKelurahan       string `json:"nama kelurahan"`
	KodeRw              string `json:"kode RW"`
	KodeRt              string `json:"kode RT"`
	NamaKepalaKeluarga  string `json:"nama kepala keluarga"`
	Baduta              string `json:"memiliki baduta"`
	Balita              string `json:"memiliki balita"`
	Pus                 string `json:"terdapat PUS"`
	PusHamil            string `json:"terdapat PUS Hamil"`
	SumberAirLayakTidak string `json:"sumber air tidak layak"`
	JambanLayakTidak    string `json:"jamban layak tidak"`
	TerlaluMuda         string `json:"terlalu muda"`
	TerlaluTua          string `json:"terlalu tua"`
	TerlaluDekat        string `json:"terlalu dekat"`
	TerlaluBanyak       string `json:"terlalu banyak"`
	RisikoStunting      string `json:"beresiko stunting"`
}
