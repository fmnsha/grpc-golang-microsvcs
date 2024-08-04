package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"users-microservice/pkg/util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Mdb *mongo.Client

func InitDB() {
	dbFullHost := util.GetEnv("DB_FULL_HOST", "")
	dbPass := util.GetEnv("DB_PASSWORD", "")
	dbUser := util.GetEnv("DB_USER", "")
	host := util.GetEnv("DB_HOST", "localhost")
	dbPort := util.GetEnv("DB_PORT", "27017")

	var uri string
	if dbFullHost != "" {
		uri = fmt.Sprintf("mongodb://%v", dbFullHost)
	} else {
		if dbPass != "" {
			uri = fmt.Sprintf(`mongodb://%v:%v@%v:%v`, dbUser, dbPass, host, dbPort)
		} else {
			uri = fmt.Sprintf("mongodb://%v:%v", host, dbPort)
		}

	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb successfully connected and pinged.")

	Mdb = client

}
