package dbops

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)
var dbClient *mongo.Client
var err error

func init() {
	// 注意格式
	dbClient, err = mongo.Connect(context.TODO(), "mongodb://mainzxq:mainzxq@10.211.55.5:27927")
	if err != nil {
		log.Fatal(err)
	}
}