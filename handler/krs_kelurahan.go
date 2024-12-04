package handler

import (
	krs_kel "krs-agg/krs_kelurahan"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type krsKelHandler struct {
	krsService krs_kel.Service
}

func NewKrsKelHandler(krsService krs_kel.Service) *krsKelHandler {
	return &krsKelHandler{krsService}
}

func (handler *krsKelHandler) GetKrsAllByKel(c *gin.Context) {
	periode := convertKelToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsAllByKel(periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krsesKelResponse []krs_kel.KrsKelResponse

	for _, f := range krses {
		krsKelResponse := convertToKrsKelResponse(f)
		krsesKelResponse = append(krsesKelResponse, krsKelResponse)
	}

	c.JSON(200, gin.H{
		"data": krsesKelResponse,
	})
}

func (handler *krsKelHandler) GetKrsKelById(c *gin.Context) {
	kodeProvinsi := c.Param("kdprov")
	kodeKabupaten := c.Param("kdkab")
	kodeKecamatan := c.Param("kdkec")
	kodeKelurahan := c.Param("kdkel")
	periode := convertKelToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsKelById(kodeProvinsi, kodeKabupaten, kodeKecamatan, kodeKelurahan, periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krsesKelResponse []krs_kel.KrsKelResponse

	for _, f := range krses {
		krsKelResponse := convertToKrsKelResponse(f)
		krsesKelResponse = append(krsesKelResponse, krsKelResponse)
	}

	c.JSON(200, gin.H{
		"data": krsesKelResponse,
	})
}

// RekapKrs godoc
// @Summary Get Aggregate Keluarga Beresiko Stunting By Kelurahan
// @Description Get Aggregate Keluarga Beresiko Stunting By Kelurahan
// @ID get-krs-by-kelurahan
// @Tags Get Aggregate Keluarga Beresiko Stunting By Kelurahan
// @Accept */*
// @Produce json
// @Success      200  {object}  krs_kel.Krs
// @Param kdprov query string true "Province Code"
// @Param kdkab query string true "Kabupaten Code"
// @Param kdkec query string true "Kecamatan Code"
// @Router /krsbykec [get]
func (handler *krsKelHandler) GetKrsByKec(c *gin.Context) {

	kodeProvinsi := c.Param("kdprov")
	kodeKabupaten := c.Param("kdkab")
	kodeKecamatan := c.Param("kdkec")
	periode := convertKelToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsByKec(kodeProvinsi, kodeKabupaten, kodeKecamatan, periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}
	var krsesKelResponse []krs_kel.KrsKelResponse
	for _, f := range krses {
		krsKelResponse := convertToKrsKelResponse(f)
		krsesKelResponse = append(krsesKelResponse, krsKelResponse)
	}
	c.JSON(200, gin.H{
		"data": krsesKelResponse,
	})
}

func (handler *krsKelHandler) GetKrsByKab(c *gin.Context) {

	kodeProvinsi := c.Param("kdprov")
	kodeKabupaten := c.Param("kdkab")
	periode := convertKelToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsByKab(kodeProvinsi, kodeKabupaten, periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}
	var krsesKelResponse []krs_kel.KrsKelResponse
	for _, f := range krses {
		krsKelResponse := convertToKrsKelResponse(f)
		krsesKelResponse = append(krsesKelResponse, krsKelResponse)
	}
	c.JSON(200, gin.H{
		"data": krsesKelResponse,
	})
}

func (handler *krsKelHandler) GetKrsByProv(c *gin.Context) {

	kodeProvinsi := c.Param("kdprov")
	periode := convertKelToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsByProv(kodeProvinsi, periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}
	var krsesKelResponse []krs_kel.KrsKelResponse
	for _, f := range krses {
		krsKelResponse := convertToKrsKelResponse(f)
		krsesKelResponse = append(krsesKelResponse, krsKelResponse)
	}
	c.JSON(200, gin.H{
		"data": krsesKelResponse,
	})
}

func convertKelToInts(str string, defaultValue int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}
func convertToKrsKelResponse(f krs_kel.KrsKel) krs_kel.KrsKelResponse {
	return krs_kel.KrsKelResponse{
		Periode:                             int(f.Periode),
		IdKelurahan:                         int(f.IdKelurahan),
		KodeKelurahan:                       f.KodeKelurahan,
		NamaKelurahan:                       f.NamaKelurahan,
		IdKecamatan:                         int(f.IdKecamatan),
		KodeKecamatan:                       f.KodeKecamatan,
		NamaKecamatan:                       f.NamaKecamatan,
		IdKabupaten:                         int(f.IdKabupaten),
		KodeKabupaten:                       f.KodeKabupaten,
		NamaKabupaten:                       f.NamaKabupaten,
		IdProvinsi:                          int(f.IdProvinsi),
		KodeProvinsi:                        f.KodeProvinsi,
		NamaProvinsi:                        f.NamaProvinsi,
		JumlahKeluarga:                      int(f.JumlahKeluarga),
		JumlahKeluargaSasaran:               int(f.JumlahKeluargaSasaran),
		Prioritas1:                          int(f.Prioritas1),
		Prioritas2:                          int(f.Prioritas2),
		Prioritas3:                          int(f.Prioritas3),
		Prioritas4:                          int(f.Prioritas4),
		PeringkatKesejahteraanDiatas4:       int(f.PeringkatKesejahteraanDiatas4),
		JumlahKeluargaBeresikoStunting:      int(f.JumlahKeluargaBeresikoStunting),
		JumlahPus:                           int(f.JumlahPus),
		JumlahPusHamil:                      int(f.JumlahPusHamil),
		JumlahBalita:                        int(f.JumlahBalita),
		JumlahBaduta:                        int(f.JumlahBaduta),
		JumlahBukanPesertaKbModern:          int(f.JumlahBukanPesertaKbModern),
		JumlahJambanTidakLayak:              int(f.JumlahJambanTidakLayak),
		JumlahAirTidakLayak:                 int(f.JumlahAirTidakLayak),
		JumlahTerlaluBanyak:                 int(f.JumlahTerlaluBanyak),
		JumlahTerlaluMuda:                   int(f.JumlahTerlaluMuda),
		JumlahTerlaluTua:                    int(f.JumlahTerlaluTua),
		JumlahTerlaluDekat:                  int(f.JumlahTerlaluDekat),
		JumlahKeluargaTidakBeresikoStunting: int(f.JumlahKeluargaTidakBeresikoStunting),
	}
}
