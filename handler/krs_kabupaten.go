package handler

import (
	krs_kab "krs-agg/krs_kabupaten"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type krsKabHandler struct {
	krsService krs_kab.Service
}

func NewKrsKabHandler(krsService krs_kab.Service) *krsKabHandler {
	return &krsKabHandler{krsService}
}

func (handler *krsKabHandler) GetKrsAllByKab(c *gin.Context) {
	periode := convertKecToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsAllByKab(periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krsesKabResponse []krs_kab.KrsKabResponse

	for _, f := range krses {
		krsKabResponse := convertToKrsKabResponse(f)
		krsesKabResponse = append(krsesKabResponse, krsKabResponse)
	}

	c.JSON(200, gin.H{
		"data": krsesKabResponse,
	})
}

func (handler *krsKabHandler) GetKrsKabById(c *gin.Context) {
	kodeProvinsi := c.Param("kdprov")
	kodeKabupaten := c.Param("kdkab")
	periode := convertKabToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsKabById(kodeProvinsi, kodeKabupaten, periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krsesKabResponse []krs_kab.KrsKabResponse

	for _, f := range krses {
		krsKabResponse := convertToKrsKabResponse(f)
		krsesKabResponse = append(krsesKabResponse, krsKabResponse)
	}

	c.JSON(200, gin.H{
		"data": krsesKabResponse,
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
func (handler *krsKabHandler) GetKrsByProv(c *gin.Context) {

	kodeProvinsi := c.Param("kdprov")
	periode := convertKecToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsByProv(kodeProvinsi, periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}
	var krsesKabResponse []krs_kab.KrsKabResponse
	for _, f := range krses {
		krsKabResponse := convertToKrsKabResponse(f)
		krsesKabResponse = append(krsesKabResponse, krsKabResponse)
	}
	c.JSON(200, gin.H{
		"data": krsesKabResponse,
	})
}

// func (handler *krsKecHandler) GetKrsByKab(c *gin.Context) {

// 	kodeProvinsi := c.Param("kdprov")
// 	kodeKabupaten := c.Param("kdkab")
// 	periode := convertKecToInts(c.Param("periode"), 0)
// 	krses, err := handler.krsService.KrsByKab(kodeProvinsi, kodeKabupaten, periode)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors message": err})
// 		return
// 	}
// 	var krsesKecResponse []krs_kec.KrsKecResponse
// 	for _, f := range krses {
// 		krsKecResponse := convertToKrsKecResponse(f)
// 		krsesKecResponse = append(krsesKecResponse, krsKecResponse)
// 	}
// 	c.JSON(200, gin.H{
// 		"data": krsesKecResponse,
// 	})
// }

func convertKabToInts(str string, defaultValue int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}
func convertToKrsKabResponse(f krs_kab.KrsKab) krs_kab.KrsKabResponse {
	return krs_kab.KrsKabResponse{
		Periode:                             int(f.Periode),
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
