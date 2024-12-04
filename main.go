package main

import (
	"fmt"
	"krs-agg/handler"
	krs_kabupaten "krs-agg/krs_kabupaten"
	krs_kecamatan "krs-agg/krs_kecamatan"
	krs_kelurahan "krs-agg/krs_kelurahan"
	krs_provinsi "krs-agg/krs_provinsi"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	pagination "github.com/webstradev/gin-pagination"

	_ "krs-agg/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// swagger embed files
// @title BKKBN Digital Service - Rekapitulasi Keluarga Beresiko Stunting
// @version         2.1
// @description     Digital Service BKKBN for Integration
// @termsOfService  https://bkkbn.go.id
// @contact.name   Direktorat Teknologi Informasi dan Data
// @contact.url    https://bkkbn.go.id
// @contact.email  dittifdok@bkkbn.go.id
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      gorest-2022-krs-ds-bkkbn-2022-krs-agg.apps.ocp-dev.bkkbn.go.id
// @BasePath  /v1/api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_database := os.Getenv("DB_DATABASE")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", db_host, db_username, db_password, db_database, db_port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sch_bkkbn.", // schema name
			SingularTable: false,
		}, PrepareStmt: true,
	})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	//------------------------ Start - Handler Kelurahan ------------------------//
	krsKelRepository := krs_kelurahan.NewRepository(db)
	krsKelService := krs_kelurahan.NewService(krsKelRepository)
	krsKelHandler := handler.NewKrsKelHandler(krsKelService)
	//------------------------ End - Handler Kelurahan ------------------------//

	//------------------------ Start - Handler Kecamatan ------------------------//
	krsKecRepository := krs_kecamatan.NewRepository(db)
	krsKecService := krs_kecamatan.NewService(krsKecRepository)
	krsKecHandler := handler.NewKrsKecHandler(krsKecService)
	//------------------------ End - Handler Kecamatan ------------------------//

	//------------------------ Start - Handler Kabupaten ------------------------//
	krsKabRepository := krs_kabupaten.NewRepository(db)
	krsKabService := krs_kabupaten.NewService(krsKabRepository)
	krsKabHandler := handler.NewKrsKabHandler(krsKabService)
	//------------------------ Start - Handler Kabupaten ------------------------//

	//------------------------ Start - Handler Provinsi ------------------------//
	krsProvRepository := krs_provinsi.NewRepository(db)
	krsProvService := krs_provinsi.NewService(krsProvRepository)
	krsProvHandler := handler.NewKrsProvHandler(krsProvService)
	//------------------------ Start - Handler Provinsi ------------------------//

	fmt.Println("================== Start Engine for KRS-AGG -- V4 ==================")

	r := gin.Default()

	corsConfig := cors.DefaultConfig()

	// corsConfig.AllowOrigins = []string{"https://example.com"}
	corsConfig.AllowAllOrigins = true
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))

	r.Use(pagination.Default())

	v4 := r.Group("/v4/api")

	//------------------------ Start - Level Kelurahan ------------------------//
	v4.GET("/rekapkrs/all/kelurahan/:periode", krsKelHandler.GetKrsAllByKel)
	v4.GET("/rekapkrs/kel/provinsi/:periode/:kdprov", krsKelHandler.GetKrsByProv)
	v4.GET("/rekapkrs/kel/kabupaten/:periode/:kdprov/:kdkab", krsKelHandler.GetKrsByKab)
	v4.GET("/rekapkrs/kel/kecamatan/:periode/:kdprov/:kdkab/:kdkec", krsKelHandler.GetKrsByKec)
	v4.GET("/rekapkrs/kel/kelurahan/:periode/:kdprov/:kdkab/:kdkec/:kdkel", krsKelHandler.GetKrsKelById)
	//------------------------ End - Level Kelurahan ------------------------//

	//------------------------ Start - Level Kecamatan ------------------------//
	v4.GET("/rekapkrs/all/kecamatan/:periode", krsKecHandler.GetKrsAllByKec)
	v4.GET("/rekapkrs/kec/kabupaten/:periode/:kdprov/:kdkab", krsKecHandler.GetKrsByKab)
	v4.GET("/rekapkrs/kec/kecamatan/:periode/:kdprov/:kdkab/:kdkec", krsKecHandler.GetKrsKecById)
	//------------------------ End - Level Kecamatan ------------------------//

	//------------------------ Start - Level Kabupaten ------------------------//
	v4.GET("/rekapkrs/all/kabupaten/:periode", krsKabHandler.GetKrsAllByKab)
	v4.GET("/rekapkrs/kab/kabupaten/:periode/:kdprov", krsKabHandler.GetKrsByProv)
	v4.GET("/rekapkrs/kab/kabupaten/:periode/:kdprov/:kdkab", krsKabHandler.GetKrsKabById)
	//------------------------ End - Level Kabupaten ------------------------//

	//------------------------ Start - Level Provinsi ------------------------//
	v4.GET("/rekapkrs/all/provinsi/:periode", krsProvHandler.GetKrsAllByProv)
	v4.GET("/rekapkrs/prov/provinsi/:periode", krsProvHandler.GetKrsByNas)
	v4.GET("/rekapkrs/prov/provinsi/:periode/:kdprov", krsProvHandler.GetKrsProvById)
	//------------------------ End - Level Provinsi ------------------------//

	//------------------------ Start - Swagger ------------------------//
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//------------------------ End - Swagger ------------------------//

	r.Run(":8080")
}
