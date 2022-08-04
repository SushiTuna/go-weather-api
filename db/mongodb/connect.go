package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Credentials struct {
	Uri string `json:"uri"`
}

func (credential Credentials) Connect() (*mongo.Client, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	context := context.TODO()

	clientOptions := options.Client().
		ApplyURI(credential.Uri).
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(context, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client, err
}

func Setup() (*mongo.Client, error) {
	credential := LoadConfig("./conf/mongodb_dev.json")
	connection, err := credential.Connect()

	return connection, err
}

func LoadConfig(file string) Credentials {
	var credential Credentials

	configFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&credential)

	return credential
}
