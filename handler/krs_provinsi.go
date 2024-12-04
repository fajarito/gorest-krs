package handler

import (
	krs_prov "krs-agg/krs_provinsi"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type krsProvHandler struct {
	krsService krs_prov.Service
}

func NewKrsProvHandler(krsService krs_prov.Service) *krsProvHandler {
	return &krsProvHandler{krsService}
}

func (handler *krsProvHandler) GetKrsAllByProv(c *gin.Context) {
	periode := convertKecToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsAllByProv(periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krsesProvResponse []krs_prov.KrsProvResponse

	for _, f := range krses {
		krsProvResponse := convertToKrsProvResponse(f)
		krsesProvResponse = append(krsesProvResponse, krsProvResponse)
	}

	c.JSON(200, gin.H{
		"data": krsesProvResponse,
	})
}

func (handler *krsProvHandler) GetKrsProvById(c *gin.Context) {
	kodeProvinsi := c.Param("kdprov")
	periode := convertProvToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsProvById(kodeProvinsi, periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krsesProvResponse []krs_prov.KrsProvResponse

	for _, f := range krses {
		krsProvResponse := convertToKrsProvResponse(f)
		krsesProvResponse = append(krsesProvResponse, krsProvResponse)
	}

	c.JSON(200, gin.H{
		"data": krsesProvResponse,
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
// func (handler *krsKabHandler) GetKrsByProv(c *gin.Context) {

// 	kodeProvinsi := c.Param("kdprov")
// 	periode := convertKecToInts(c.Param("periode"), 0)
// 	krses, err := handler.krsService.KrsByProv(kodeProvinsi, periode)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors message": err})
// 		return
// 	}
// 	var krsesProvResponse []krs_prov.KrsKabResponse
// 	for _, f := range krses {
// 		krsProvResponse := convertToKrsProvResponse(f)
// 		krsesProvResponse = append(krsesProvResponse, krsProvResponse)
// 	}
// 	c.JSON(200, gin.H{
// 		"data": krsesProvResponse,
// 	})
// }

func (handler *krsProvHandler) GetKrsByNas(c *gin.Context) {

	periode := convertKecToInts(c.Param("periode"), 0)
	krses, err := handler.krsService.KrsByNas(periode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}
	var krsesProvResponse []krs_prov.KrsProvResponse

	for _, f := range krses {
		krsProvResponse := convertToKrsProvResponse(f)
		krsesProvResponse = append(krsesProvResponse, krsProvResponse)
	}

	c.JSON(200, gin.H{
		"data": krsesProvResponse,
	})
}

func convertProvToInts(str string, defaultValue int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}
func convertToKrsProvResponse(f krs_prov.KrsProv) krs_prov.KrsProvResponse {
	return krs_prov.KrsProvResponse{
		Periode:                             int(f.Periode),
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
