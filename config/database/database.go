package database

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("FAIL TO LOAD CONFIG VALUE")
	}

	port := viper.Get("database.port").(int)
	//username := viper.Get("database.username").(string)
	//password := viper.Get("database.password").(string)

	fmt.Println(port)

	clientOptions := options.Client().ApplyURI("localhost:" + string(port))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if client != nil {
		panic("FAIL TO CONNECT TO DB")
	}
}
