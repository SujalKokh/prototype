package config

import (
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"prototype/routes"
)

func InitServer(){
	// db := InitDatabase()
	r := gin.Default()

	// routes.InitializeRoutes(r)
	routes.InitRoutes(r)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		r.Run()
	} else {
		r.Run(os.Getenv("SERVER_PORT"))
	}
}