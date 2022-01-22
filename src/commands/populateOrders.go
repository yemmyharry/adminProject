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
		var orderItems []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {
			orderItems = append(orderItems, models.OrderItem{
				ProductName:       faker.Word(),
				Price:             float64(rand.Intn(100)),
				Quantity:          uint(rand.Intn(5)),
				AdminRevenue:      0.9 * models.OrderItem{}.Price * float64(models.OrderItem{}.Quantity),
				AmbassadorRevenue: 0.1 * models.OrderItem{}.Price * float64(models.OrderItem{}.Quantity),
			})
		}

		database.DB.Create(&models.Order{
			UserID:          rand.Intn(30),
			Code:            faker.Username(),
			AmbassadorEmail: faker.Email(),
			FirstName:       faker.FirstName(),
			LastName:        faker.LastName(),
			Email:           faker.Email(),
			Complete:        true,
			OrderItems:      orderItems,
		})
	}
}
