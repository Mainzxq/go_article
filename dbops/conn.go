package dbops

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"
)
var dbClient *mongo.Client
var err error
var m = map[string]string{"SERVICE_NAME": "mongodb"}
var cd = options.Credential{
	"SCRAM-SHA-1",
	m,
	"mainzxq",
	"mainzxq",
	"mainzxq",
}

func init() {
	// 注意格式, 新版本mongo需要使用SCRAM-SHA-1的验证方式，因此需要做上面一堆blablabla的事情
	var dbOption options.ClientOptions
	dbClient, err = mongo.Connect(context.TODO(),"mongodb://mainzxq:mainzxq@192.168.43.2:27927", dbOption.SetAuth(cd) )
	if err != nil {
		fmt.Println("database going to wrong!")
		log.Fatal(err)
	}
}