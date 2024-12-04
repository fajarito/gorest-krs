package krs_kabupaten

import "encoding/json"

type KrsInput struct {
	Nama   string      `json:"nama" binding:"required"`
	Kode   json.Number `json:"kode" binding:"required,number"`
	Alamat string      `json:"alamat" binding:"required"`
}
