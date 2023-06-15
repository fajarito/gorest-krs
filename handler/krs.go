package handler

import (
	krs_kec "gorest-krs/krs_kec"

	// "gorest/faskes"
	"net/http"

	// _ "gorest/docs"

	"github.com/gin-gonic/gin"
)

type krsHandler struct {
	krsService krs_kec.Service
}

func NewKrsHandler(krsService krs_kec.Service) *krsHandler {
	return &krsHandler{krsService}
}

// RekapKrs godoc
// @Summary Get Aggregate Keluarga Beresiko Stunting By Kelurahan
// @Description Get Aggregate Keluarga Beresiko Stunting By Kelurahan
// @ID get-krs-by-kelurahan
// @Tags Get Aggregate Keluarga Beresiko Stunting By Kelurahan
// @Accept */*
// @Produce json
// @Success      200  {object}  krs_kec.Krs
// @Param kdprov query string true "Province Code"
// @Param kdkab query string true "Kabupaten Code"
// @Param kdkec query string true "Kecamatan Code"
// @Router /krsbykec [get]
func (handler *krsHandler) GetKrsByKec(c *gin.Context) {
	// kodeProvinsi := c.Param("idprov")
	kodeProvinsi := c.Query("kdprov")
	// id := c.Param("id")
	// idprov, _ := strconv.Atoi(idProvString)

	// kodeKabupaten := c.Param("idkab")
	kodeKabupaten := c.Query("kdkab")
	// id := c.Param("id")
	// idkab, _ := strconv.Atoi(idKabString)

	// kodeKecamatan := c.Param("idkec")
	kodeKecamatan := c.Query("kdkec")
	// id := c.Param("id")
	// idkec, _ := strconv.Atoi(idKecString)

	krses, err := handler.krsService.KrsByKec(kodeProvinsi, kodeKabupaten, kodeKecamatan)
	// fa, err := handler.faskesService.FindByProv(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krsesResponse []krs_kec.KrsResponse

	for _, f := range krses {
		krsResponse := convertToKrsResponse(f)
		krsesResponse = append(krsesResponse, krsResponse)
	}

	c.JSON(200, gin.H{
		"data": krsesResponse,
	})
}

func convertToKrsResponse(f krs_kec.Krs) krs_kec.KrsResponse {
	return krs_kec.KrsResponse{

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
		JumlahKeluargaTidakBeresikoStunting: int(f.JumlahKeluargaTidakBeresikoStunting),
	}
}
