package main

import (
	"api-crud-golang/config"
	"api-crud-golang/controller"
	"api-crud-golang/model"
	"api-crud-golang/service"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/logrusorgru/aurora"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_ = godotenv.Load(".env")

	fmt.Println("Using timezone:", aurora.Green(time.Now().Location().String()))
	config.InitConfig()
}

// @title       Example API
// @version     1.0
// @description This is a sample server
// @host        localhost:7391
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath    /
func main() {
	// fmt.Println("Hello World!")

	cfg := &config.PostgresConnection{
		Host:     config.POSTGRE_HOST,
		Port:     config.POSTGRE_PORT,
		Name:     config.POSTGRE_DB_NAME,
		Username: config.POSTGRE_USERNAME,
		Password: config.POSTGRE_PASSWORD,
	}

	conn := config.NewPostgresConnection(cfg)

	db, err := conn.InitConnection()
	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(model.MainModel{})
	db.AutoMigrate(model.User{})

	service := service.NewMainService(db)
	control := controller.NewMainController(service)

	router := gin.Default()
	controller.InitRouter(router, control)

	router.Run(":7391")
}
