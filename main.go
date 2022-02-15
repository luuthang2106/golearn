package main

import (
	"context"
	"golearn/db"
	"golearn/model"
)

func main() {
	client := db.GetMongoClient("mongodb://localhost:27017")
	dbInstance := client.Database("sample")
	db.AccountDB.Init(dbInstance)

	_, _ = db.AccountDB.InsertOne(&model.Account{Username: "thang", Email: "afasfasf"})
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
