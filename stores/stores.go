package stores

import (
	"context"
	"fmt"
	"mongo-with-golang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	recentActionsPtr *mongo.Collection
	userPtr          *mongo.Collection
}

const uri = "mongodb://localhost:27017/"

func (m *MongoStore) OpenConnectionWithMongoDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{}}).Decode(&result); err != nil {
		panic(err)
	}

	m.recentActionsPtr = client.Database("Codeforces").Collection("recentactions")
	fmt.Println("Pinged Your Deployment")

}

func (m *MongoStore) StoreRecentActionsInTheDatabase(actions []models.RecentActions, err error) error {
	var toInsert []interface{}
	for _, action := range actions {
		toInsert = append(toInsert, action)
	}
	_, err1 := m.recentActionsPtr.InsertMany(context.TODO(), toInsert)
	if err1 != nil {
		fmt.Println("Error", err1)
		return nil
	}
	return nil
}

func (m *MongoStore) QueryRecentActions() ([]models.RecentActions, error) {
	//var query interface{}
	cursor, err := m.recentActionsPtr.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Error")

	}
	var results []models.RecentActions
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		fmt.Println("Erorr")
	}
	return results, err

}
