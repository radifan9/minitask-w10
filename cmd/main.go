package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/radifan9/minitask-w10/internal/configs"
	"github.com/radifan9/minitask-w10/internal/models"
	"github.com/radifan9/minitask-w10/internal/routers"
)

var users = make(map[string]models.User)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env\nCause: ", err.Error())
		return
	}

	// DB initialization
	db, err := configs.InitDB()
	if err != nil {
		log.Println("failed to connect to database\nCause: ", err.Error())
		return
	}
	defer db.Close()

	// Test DB connection
	if err := configs.TestDBCon(db); err != nil {
		log.Println("ping to DB failed\nCause: ", err.Error())
		return
	}

	// Inisiasi engine gin
	// router := gin.Default()
	router := routers.InitRouter(db)

	// Run Engine Gin
	router.Run(":8080")
}
