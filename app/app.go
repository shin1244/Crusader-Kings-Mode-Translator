package app

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 해당 콜렉션이 존재하는지 확인하는 함수
func CheckCollectionExist(collections []string, collection string) bool {
	isCollectionExist := false
	for _, c := range collections {
		if c == collection {
			isCollectionExist = true
			break
		}
	}
	return isCollectionExist
}

// 맵에서 키에 해당하는 값을 반환하는 함수 (좆같음)
func GetMapValue[K comparable, V any](m map[K]V, key K) (V, bool) {
	value, ok := m[key]
	return value, ok
}

// MongoDB에 연결하는 함수
func MongoConnect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
