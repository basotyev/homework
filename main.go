package main

import (
	"context"
	"fmt"
	"lesson13/db"
	RP "lesson13/user/repository"
)

func main() {
	postgresDb, err := db.NewPostgresConnection("postgres://postgres:pass@localhost:5439/test-db?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	mongoDb, err := db.NewMongoConnection("mongodb://root:example@localhost:27017")
	if err != nil {
		fmt.Println(err)
	}
	userPostgresRepo := RP.NewPostgresUserRepository(postgresDb)
	userMongoRepo := RP.NewMongoRepository(mongoDb)

	newUser, err := userPostgresRepo.GetUserById(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("New user from: %v\n", *newUser)
	userFromMongo, err := userMongoRepo.GetUserById(context.Background(), newUser.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*userFromMongo)
	// TODO 1.Реализовать репозиторий для сущности Task
	// TODO 2. Написать миграции для Task
	// TODO 3.В main.go добавить примеры использования CRUD операций
}

//var aza models.User
//aza = models.User{
//	Name:  "Azamat",
//	Email: "azamat@example.com",
//	Age:   30,
//}
//err = userPostgresRepo.CreateUser(context.Background(), &aza)
//if err != nil {
//	fmt.Println(err)
//	return
//}
//fmt.Printf("First user: %v\n", &aza)
//aza.Name = "Azamat 2.0"
//err = userPostgresRepo.CreateUser(context.Background(), &aza)
//if err != nil {
//	fmt.Println(err)
//	return
//}
//fmt.Printf("Second user: %v\n", &aza)
//var Beka models.User
//Beka = models.User{
//	Id:    aza.Id,
//	Name:  "Beka",
//	Email: "beka@example.com",
//	Age:   20,
//}
//err = userPostgresRepo.UpdateUserById(context.Background(), &Beka)
//if err != nil {
//	fmt.Println(err)
//	return
//}
//fmt.Printf("Second user after update: %v\n", &Beka)
//err = userPostgresRepo.RemoveUserById(context.Background(), 1)
//if err != nil {
//	fmt.Println(err)
//	return
//}
