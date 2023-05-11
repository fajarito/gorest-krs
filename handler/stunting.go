package handler

import (
	"fmt"
	"gorest-pk/stunting"

	// "gorest/faskes"
	"net/http"

	// _ "gorest/docs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type stuntingHandler struct {
	stuntingService stunting.Service
}

func NewStuntingHandler(stuntingService stunting.Service) *stuntingHandler {
	return &stuntingHandler{stuntingService}
}

// GetStunting godoc
// @Summary Get List Detil Keluarga
// @Description Get List Detil Keluarga
// @ID get-stunting
// @Tags Get List Detil Keluarga
// @Accept */*
// @Produce json
// @Success      200  {object}  stunting.Stunting
// @Param nik path string true "nik"
// @Router /cekkeluargadetail/{nik} [get]
func (handler *stuntingHandler) GetStunting(c *gin.Context) {

	id := c.Param("id")

	fa, err := handler.stuntingService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data Tidak Ditemukan / NIK Harus 16 Digit / NIK Harus terdiri dari Numeric"})
		return
	}

	stuntingResponse := convertToStuntingResponse(fa)

	c.JSON(200, gin.H{
		"Status":  "Success",
		"Message": "Data Ditemukan",
		"data":    stuntingResponse,
	})
}

// GetKeluargaPusHamil godoc
// @Summary Get List Detil Keluarga (PUS) yang sedang dalam kondisi Hamil
// @Description Get List Detil Keluarga (PUS) yang sedang dalam kondisi Hamil
// @ID get-keluarga-pus-hamil
// @Tags Get List Detil Keluarga (PUS) yang sedang dalam kondisi Hamil
// @Accept */*
// @Produce json
// @Success      200  {object}  stunting.Stunting
// @Param nik path string true "nik"
// @Router /cekpushamil/{nik} [get]
func (handler *stuntingHandler) GetKeluargaPusHamil(c *gin.Context) {

	id := c.Param("id")

	fa, err := handler.stuntingService.FindByPusHamil(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data Tidak Ditemukan / NIK Harus 16 Digit / NIK Harus terdiri dari Numeric"})
		return
	}

	stuntingResponse := convertToStuntingResponse(fa)

	c.JSON(200, gin.H{
		"Status":  "Success",
		"Message": "Data Keluarga (Memiliki PUS Hamil) Ditemukan",
		"data":    stuntingResponse,
	})
}

// GetKeluargaResikoStunting godoc
// @Summary Get List Detil Keluarga yang beresiko Stunting
// @Description Get List Detil Keluarga yang beresiko Stunting
// @ID get-keluarga-resiko-stunting
// @Tags Get List Detil Keluarga yang beresiko Stunting
// @Accept */*
// @Produce json
// @Success      200  {object}  stunting.Stunting
// @Param nik path string true "nik"
// @Router /cekresikostunting/{nik} [get]
func (handler *stuntingHandler) GetKeluargaResikoStunting(c *gin.Context) {

	id := c.Param("id")

	fa, err := handler.stuntingService.FindByStunting(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data Tidak Ditemukan / NIK Harus 16 Digit / NIK Harus terdiri dari Numeric"})
		return
	}

	stuntingResponse := convertToStuntingResponse(fa)

	c.JSON(200, gin.H{
		"Status":  "Success",
		"Message": "Data Keluarga Beresiko Stunting Ditemukan",
		"data":    stuntingResponse,
	})
}

// GetKeluargaBaduta godoc
// @Summary Get List Detil Keluarga yang memiliki Baduta (Bayi Usia 2 Tahun)
// @Description Get List Detil Keluarga yang memiliki Baduta (Bayi Usia 2 Tahun)
// @ID get-keluarga-baduta
// @Tags Get List Detil Keluarga yang memiliki Baduta (Bayi Usia 2 Tahun)
// @Accept */*
// @Produce json
// @Success      200  {object}  stunting.Stunting
// @Param nik path string true "nik"
// @Router /cekbaduta/{nik} [get]
func (handler *stuntingHandler) GetKeluargaBaduta(c *gin.Context) {

	id := c.Param("id")

	fa, err := handler.stuntingService.FindByBaduta(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data Tidak Ditemukan / NIK Harus 16 Digit / NIK Harus terdiri dari Numeric"})
		return
	}

	stuntingResponse := convertToStuntingResponse(fa)

	c.JSON(200, gin.H{
		"Status":  "Success",
		"Message": "Data Keluarga Beresiko Stunting Ditemukan",
		"data":    stuntingResponse,
	})
}

func (handler *stuntingHandler) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}

func (handler *stuntingHandler) WhoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "It's Me",
	})
}

func (handler *stuntingHandler) StuntingHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(200, gin.H{"id": id, "title": title})
}

func (handler *stuntingHandler) QueryHandler(c *gin.Context) {
	nama := c.Query("nama")
	kode := c.Query("kode")
	c.JSON(200, gin.H{"nama": nama, "kode": kode})
}

func (handler *stuntingHandler) PostfaskesHandler(c *gin.Context) {
	var stuntingInput stunting.StuntingInput

	err := c.ShouldBindJSON(&stuntingInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s,condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": errorMessages,
		})
		return

	}
	c.JSON(200, gin.H{
		"nama":   stuntingInput.Nama,
		"kode":   stuntingInput.Kode,
		"alamat": stuntingInput.Alamat,
	})
}

func convertToStuntingResponse(f stunting.Stunting) stunting.StuntingResponse {
	return stunting.StuntingResponse{
		Nik:                 f.Nik,
		KodeProvinsi:        f.KodeProvinsi,
		NamaProvinsi:        f.NamaProvinsi,
		KodeKabupaten:       f.KodeKabupaten,
		NamaKabupaten:       f.NamaKabupaten,
		KodeKecamatan:       f.KodeKecamatan,
		NamaKecamatan:       f.NamaKecamatan,
		KodeKelurahan:       f.KodeKelurahan,
		NamaKelurahan:       f.NamaKelurahan,
		KodeRw:              f.KodeRw,
		KodeRt:              f.KodeRt,
		NamaKepalaKeluarga:  f.NamaKepalaKeluarga,
		Baduta:              f.Baduta,
		Balita:              f.Balita,
		Pus:                 f.Pus,
		PusHamil:            f.PusHamil,
		SumberAirLayakTidak: f.SumberAirLayakTidak,
		JambanLayakTidak:    f.JambanLayakTidak,
		TerlaluMuda:         f.TerlaluMuda,
		TerlaluTua:          f.TerlaluTua,
		TerlaluDekat:        f.TerlaluDekat,
		TerlaluBanyak:       f.TerlaluBanyak,
		RisikoStunting:      f.RisikoStunting,
	}
}
