package insert_speed

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
	"vs/pkg/tests"
	"vs/pkg/utils"
)

func runMongo(totalCnt, rps int) {
	// MongoDB에 연결
	clientOptions := options.Client().ApplyURI("mongodb://mongo:mongo@localhost:27017")
	clientOptions.SetMaxPoolSize(100)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// 연결 테스트
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// 데이터베이스 및 컬렉션 참조
	collection := client.Database("mongo").Collection("users")

	// 컬렉션의 모든 데이터 삭제
	_, err = collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// 호출
	start := time.Now()
	status := tests.LoopFunction(collection, totalCnt, rps, createMongo)
	fmt.Println("[ 완료 ]\n▶️ 동작 시간 : ", time.Since(start), "\n▶️ 성공 : ", status.Success, "\n▶️ 실패 : ", status.Failed)
}

func createMongo(arg interface{}) error {
	collection := arg.(*mongo.Collection)
	user := User{
		Name: utils.RandomString(30),
		Age:  rand.Intn(100) + 1,
	}
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}
