package dbops

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/mongo/driver/auth"
	"log"
)
var dbClient *mongo.Client
var err error

func init() {
	// 注意格式
	auth.SCRAMSHA1
	dbClient, err = mongo.Connect(context.TODO(), "mongodb://mainzxq:mainzxq@192.168.43.2:27927")
	if err != nil {
		log.Fatal(err)
	}
}