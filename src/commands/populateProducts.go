package main

import (
	"adminProject/src/database"
	"adminProject/src/models"
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

func main() {

	database.Connect()

	for i := 0; i < 30; i++ {
		products := models.Product{
			IDModel:     models.IDModel{},
			Name:        faker.Word(),
			Description: faker.Sentence(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 10),
		}
		database.DB.Create(&products)
	}
}
