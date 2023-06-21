package main 

import (
	"fmt"
	"log"
	"flag"

	// "go-template/db/mysql"
	"go-template/vars"
	"go-template/services/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	port := flag.String("port", "8000", "app port"); 
	flag.Parse() 
	
	SetupConfigFromEnv() // comment this if not using .env
	// mysql.InitDB() // uncomment when db is exist

	router := GetRouter() // router.go
	HttpServer(router, *port)
}

func SetupConfigFromEnv() {
	// load env file 
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// load env variables into struct 
	envsPrefix := ""
	err = envconfig.Process(envsPrefix, &vars.Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("=====> Env file setup: success")
}

func HttpServer(r *echo.Echo, port string) {
	externalIp, _ := utils.ExternalIP()
	logLine := fmt.Sprintf(
		"=====> HTTP-server on: \n http://localhost:%s \n http://%s:%s",
		port, externalIp, port,
	)
	fmt.Println(logLine)
	r.Logger.Fatal(r.Start(":" + port))
}
