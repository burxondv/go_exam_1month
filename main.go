package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	PostgresUser     = "postgres"
	PostgresPassword = "12345"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresDatabase = "go_exam"
)

func main() {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		PostgresDatabase,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	m := NewDBManager(db)

	// --- CREATE CAR AND IMAGES

	// id, err := m.CreateCar(&Car{		// First car
	// 	model:    "Lambo Urus",
	// 	color:    "Blue",
	// 	price:    450000,
	// 	year:     2022,
	// 	imageUrl: "exam_url",
	// 	images: []*CarImage{
	// 		{
	// 			imageUrl:       "exam_url1",
	// 			sequenceNumber: 1,
	// 		},
	// 		{
	// 			imageUrl:       "exam_url2",
	// 			sequenceNumber: 2,
	// 		},
	// 		{
	// 			imageUrl:       "exam_url3",
	// 			sequenceNumber: 3,
	// 		},
	// 	},
	// })
	// id, err := m.CreateCar(&Car{		// Second car
	// 	model:    "Tesla",
	// 	color:    "Purple",
	// 	price:    100000,
	// 	year:     2019,
	// 	imageUrl: "teslaX3_url",
	// 	images: []*CarImage{
	// 		{
	// 			imageUrl:       "tesl20_url1",
	// 			sequenceNumber: 1,
	// 		},
	// 		{
	// 			imageUrl:       "tesla12_url2",
	// 			sequenceNumber: 2,
	// 		},
	// 		{
	// 			imageUrl: "teslss_url3",
	// 			sequenceNumber: 3,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("failed to create car: %v", err)
	// }
	// fmt.Println(id)
	// id, err := m.CreateCar(&Car{ // Third car
	// 	model:    "Lambo Urus",
	// 	color:    "Red",
	// 	price:    400000,
	// 	year:     2018,
	// 	imageUrl: "la3_url",
	// 	images: []*CarImage{
	// 		{
	// 			imageUrl:       "lambo_url1",
	// 			sequenceNumber: 1,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("failed to create car: %v", err)
	// }
	// fmt.Println(id)


	// --- GET CAR

	// car, err := m.GetCar(2)
	// if err != nil {
	// 	log.Fatalf("failed to get car: %v", err)
	// }
	// fmt.Println(car)


	// --- GET ALL CARS

	// res, err := m.GetAllCars(&GetCarParam{
	// 	model: "Lambo Urus",
	// 	limit: 10,
	// 	page: 1,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to get all cars: %v", err)
	// }
	// fmt.Printf("%v", res)


	// --- UPDATE CARS

	// err = m.UpdateCar(&Car{
	// 	id: 4,
	// 	model: "lambo",
	// 	color: "red",
	// 	price: 20000,
	// 	year: 2022,
	// 	imageUrl: "new_url",
	// })
	// if err != nil {
	// 	log.Fatalf("failed to update car: %v", err)
	// }


	// --- DELETE CAR

	// err = m.DeleteCar(2)
	// if err != nil {
	// 	log.Fatalf("failed to delete car: %v", err)
	// }

}
