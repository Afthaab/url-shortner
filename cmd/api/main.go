package main

import (
	"fmt"
	"log"

	"github.com/afthaab/urlshortner/pkg/db"
	"github.com/afthaab/urlshortner/pkg/di"
	"github.com/afthaab/urlshortner/pkg/repository"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Getting the database address for both DB's
	rdb1 := db.CreateRedisClientZero()
	rdb2 := db.CreateRedisClientOne()

	// passing the reference to the repository
	repoInterface := repository.NewUrlRepository(rdb1, rdb2)

	server, err := di.InitializeAPI(repoInterface)
	if err != nil {
		log.Fatalln("Could not start the Server : ", err)
	} else {
		server.Start()
	}

	fmt.Println("Server Started at localhost : 3000")
}
