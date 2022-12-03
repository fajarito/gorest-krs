package stunting

import "encoding/json"

type StuntingInput struct {
	Nama   string      `json:"nama" binding:"required"`
	Kode   json.Number `json:"kode" binding:"required,number"`
	Alamat string      `json:"alamat" binding:"required"`
}
