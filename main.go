package main

import (
	"fmt"
	"gorest-pk/handler"
	"gorest-pk/stunting"
	"log"
	"os"

	"github.com/gin-contrib/cors"

	_ "gorest-pk/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// swagger embed files
// @title BKKBN Digital Service Documentation
/// @version         1.0
// @description     Digital Service BKKBN for Integration
// @termsOfService  https://bkkbn.go.id
// @contact.name   Direktorat Teknologi Informasi dan Data
// @contact.url    https://bkkbn.go.id
// @contact.email  dittifdok@bkkbn.go.id
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
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
	// dsn := "host=36.37.120.165 user=usr_bkkbn password=bkkbn$1 dbname=BKKBN port=5435 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sch_bkkbn.", // schema name
			SingularTable: false,
		}, PrepareStmt: true,
	})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	stuntingRepository := stunting.NewRepository(db)
	stuntingService := stunting.NewService(stuntingRepository)
	stuntingHandler := handler.NewStuntingHandler(stuntingService)

	fmt.Println("==================v3==================v3==================v3==================v")

	r := gin.Default()

	v1 := r.Group("/v1")
	corsConfig := cors.DefaultConfig()

	// corsConfig.AllowOrigins = []string{"https://example.com"}
	corsConfig.AllowAllOrigins = true
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))

	v1.GET("/ping", stuntingHandler.PingHandler)

	v1.GET("/who", stuntingHandler.WhoHandler)

	v1.GET("/stunting/:id/:title", stuntingHandler.StuntingHandler)

	v1.GET("query", stuntingHandler.QueryHandler)

	v1.POST("/poststunting", stuntingHandler.PostfaskesHandler)

	v2 := r.Group("/api/v1")
	v2.GET("/cekkeluargadetail/:id", stuntingHandler.GetStunting)
	v2.GET("/cekresikostunting/:id", stuntingHandler.GetKeluargaResikoStunting)
	v2.GET("/cekpushamil/:id", stuntingHandler.GetKeluargaPusHamil)
	v2.GET("/cekbaduta/:id", stuntingHandler.GetKeluargaBaduta)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
