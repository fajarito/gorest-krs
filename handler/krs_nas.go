package handler

import (
	krs_nas "gorest-krs/krs_nas"

	// "gorest/faskes"
	"net/http"

	// _ "gorest/docs"

	"github.com/gin-gonic/gin"
)

type krsNasHandler struct {
	krsNasService krs_nas.Service
}

func NewKrsNasHandler(krsNasService krs_nas.Service) *krsNasHandler {
	return &krsNasHandler{krsNasService}
}

// GetStunting godoc
// @Summary Get Aggregate Keluarga Beresiko Stunting By Province
// @Description Get Aggregate Keluarga Beresiko Stunting By Province
// @ID get-krs-by-prov
// @Tags Get Aggregate Keluarga Beresiko Stunting By Province
// @Accept */*
// @Produce json
// @Success      200  {object}  krs_nas.KrsNas
// @Router /krsbynas [get]
func (handler *krsNasHandler) GetKrsByNas(c *gin.Context) {
	// kodeProvinsi := c.Param("idprov")

	// id := c.Param("id")
	// idprov, _ := strconv.Atoi(idProvString)

	krsNases, err := handler.krsNasService.KrsByNas()
	// fa, err := handler.faskesService.FindByProv(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krsNasesResponse []krs_nas.KrsNasResponse

	for _, f := range krsNases {
		krsNasResponse := convertToKrsNasesResponse(f)
		krsNasesResponse = append(krsNasesResponse, krsNasResponse)
	}

	c.JSON(200, gin.H{
		"data": krsNasesResponse,
	})
}

func convertToKrsNasesResponse(f krs_nas.KrsNas) krs_nas.KrsNasResponse {
	return krs_nas.KrsNasResponse{

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
