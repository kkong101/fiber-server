package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kkong101/fiber-server/app/module/product/entity"
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

	collection := client.Database("urikiri").Collection("community_posts")

	oid, err := primitive.ObjectIDFromHex("6364d5ed099a885b14e6812a")
	cursor, err := collection.Find(context.TODO(), bson.M{"_id": oid})
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []*entity.Post
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results in a collection")
	doc, _ := json.MarshalIndent(results, "", "  ")
	fmt.Println(string(doc))
	//for _, result := range results {
	//	doc, _ := json.MarshalIndent(result, "", "  ")
	//	fmt.Println(string(doc))
	//}

}
