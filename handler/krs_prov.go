package handler

import (
	krs_prov "gorest-krs/krs_prov"

	// "gorest/faskes"
	"net/http"

	// _ "gorest/docs"

	"github.com/gin-gonic/gin"
)

type krsProvHandler struct {
	krsProvService krs_prov.Service
}

func NewKrsProvHandler(krsProvService krs_prov.Service) *krsProvHandler {
	return &krsProvHandler{krsProvService}
}

// RekapKrs godoc
// @Summary Get Aggregate Keluarga Beresiko Stunting By Kabupaten
// @Description Get Aggregate Keluarga Beresiko Stunting By Kabupaten
// @ID get-krs-by-kab
// @Tags Get Aggregate Keluarga Beresiko Stunting By Kabupaten
// @Accept */*
// @Produce json
// @Success      200  {object}  krs_prov.KrsProv
// @Param kdprov query string true "Province Code"
// @Router /krsbyprov [get]
func (handler *krsProvHandler) GetKrsByProv(c *gin.Context) {
	// kodeProvinsi := c.Param("idprov")
	kodeProvinsi := c.Query("kdprov")
	// id := c.Param("id")
	// idprov, _ := strconv.Atoi(idProvString)

	krsProvs, err := handler.krsProvService.KrsByProv(kodeProvinsi)
	// fa, err := handler.faskesService.FindByProv(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krsProvsResponse []krs_prov.KrsProvResponse

	for _, f := range krsProvs {
		krsProvResponse := convertToKrsProvsResponse(f)
		krsProvsResponse = append(krsProvsResponse, krsProvResponse)
	}

	c.JSON(200, gin.H{
		"data": krsProvsResponse,
	})
}

func convertToKrsProvsResponse(f krs_prov.KrsProv) krs_prov.KrsProvResponse {
	return krs_prov.KrsProvResponse{

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
		JumlahKeluargaTidakBeresikoStunting: int(f.JumlahKeluargaTidakBeresikoStunting),
	}
}
