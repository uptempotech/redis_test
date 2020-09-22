package main

import (
	"log"

	"github.com/uptempotech/redistest/global"
	"github.com/uptempotech/redistest/models"
	"github.com/uptempotech/redistest/services"
)

func main() {
	redisClient := services.NewRedisClient()

	key1 := "sampleKey"
	value1 := &models.ValueEx{Name: "someName", Email: "someemail@abc.com"}
	value2 := &models.ValueEx{}

	err := redisClient.SetKey(key1, value1, global.KeepTTL)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}

	err = redisClient.GetKey(key1, value2)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}

	log.Printf("Name: %s", value2.Name)
	log.Printf("Email: %s", value2.Email)
}
