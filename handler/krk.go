package handler

import (
	krk_kab "gorest-krs/krk_kab"
	"math"
	"strconv"

	// "gorest/faskes"
	"net/http"

	// _ "gorest/docs"

	"github.com/gin-gonic/gin"
)

type krkHandler struct {
	krkService krk_kab.Service
}

// ////////////////
type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
	TotalRows    int
}

//////////////////

func NewKrkHandler(krkService krk_kab.Service) *krkHandler {
	return &krkHandler{krkService}
}

// RekapKrs godoc
// @Summary Get Aggregate Keluarga Beresiko Stunting By Kecamatan
// @Description Get Aggregate Keluarga Beresiko Stunting By Kecamatan
// @ID get-krs-by-kec
// @Tags Get Aggregate Keluarga Beresiko Stunting By Kecamatan
// @Accept */*
// @Produce json
// @Success      200  {object}  krk_kab.Krk
// @Param kdprov query string true "Province Code"
// @Param kdkab query string true "Kabupaten Code"
// @Router /krsbykab [get]
func (handler *krkHandler) GetKrkByKab(c *gin.Context) {

	kodeProvinsi := c.Query("kdprov")
	kodeKabupaten := c.Query("kdkab")

	page := convertToInt(c.DefaultQuery("page", "1"), 1)
	pageSize := convertToInt(c.DefaultQuery("pageSize", "10"), 10)

	if pageSize <= 0 || pageSize > 15 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid pageSize value. Must be between 1 and 100",
		})
		return
	}

	krkes, err := handler.krkService.KrkByKab(kodeProvinsi, kodeKabupaten, page, pageSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	totalItems, err := handler.krkService.GetTotalKrkCount(kodeProvinsi, kodeKabupaten)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var krkesResponse []krk_kab.KrkkabResponse

	for _, f := range krkes {
		krkResponse := convertToKrkResponse(f)
		krkesResponse = append(krkesResponse, krkResponse)
	}
	response := gin.H{
		"data": krkesResponse,
	}

	if totalItems > 0 {
		totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
		response["pagination"] = gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalPages":  totalPages,
			"totalItems":  totalItems,
		}
	}
	c.JSON(http.StatusOK, response)

}

func (handler *krkHandler) GetKrkByKecDetail(c *gin.Context) {

	kodeProvinsi := c.Query("kdprov")
	kodeKabupaten := c.Query("kdkab")
	kodeKecamatan := c.Query("kdkec")

	// page := convertToInt(c.DefaultQuery("page", "1"), 1)
	// pageSize := convertToInt(c.DefaultQuery("pageSize", "10"), 10)

	// if pageSize <= 0 || pageSize > 15 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Invalid pageSize value. Must be between 1 and 100",
	// 	})
	// 	return
	// }

	// krkes, err := handler.krkService.KrkByKecDetail(kodeProvinsi, kodeKabupaten, kodeKecamatan, page, pageSize)
	krkes, err := handler.krkService.KrkByKecDetail(kodeProvinsi, kodeKabupaten, kodeKecamatan)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors message": err})
		return
	}

	var krkesResponse []krk_kab.KrkkabResponse

	for _, f := range krkes {
		krkResponse := convertToKrkResponse(f)
		krkesResponse = append(krkesResponse, krkResponse)
	}
	response := gin.H{
		"data": krkesResponse,
	}

	c.JSON(http.StatusOK, response)

}

func convertToInt(str string, defaultValue int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}

func convertToKrkResponse(f krk_kab.Krk) krk_kab.KrkkabResponse {
	return krk_kab.KrkkabResponse{

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
