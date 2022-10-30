package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

// main StartApplication
func main() {
	// Dependency Injection
	fx.New()

	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("FAIL TO LOAD CONFIG VALUE")
	}

	//port := viper.Get("database.port").(int)
	//username := viper.Get("database.username").(string)
	//password := viper.Get("database.password").(string)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username: "root",
		Password: "example",
	})
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// https://github.com/kyriediculous/go-grpc-mongodb/blob/master/server/main.go

	collection := client.Database("urikiri").Collection("products")

	oid, err := primitive.ObjectIDFromHex("634590139949ea74c8f42465")
	res, err := collection.Find(context.TODO(), bson.M{"_id": oid})

	for res.Next(context.TODO()) {
		fmt.Println(res)
	}
}
