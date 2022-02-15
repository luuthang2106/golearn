package db

import (
	"context"
	"golearn/model"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var AccountDB = Model{collName: "account"}

type Model struct {
	collName string
	collIns  *mongo.Collection
	template model.Account
}

func GetMongoClient(uri string) *mongo.Client {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}
func (m *Model) Init(db *mongo.Database) {
	m.collIns = db.Collection(m.collName)
}
func (m *Model) FindOne(filter interface{}) (interface{}, error) {
	t := reflect.TypeOf(&m.template).Elem()
	v := reflect.New(t).Interface()
	if err := m.collIns.FindOne(context.TODO(), filter).Decode(v); err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Model) InsertOne(entity interface{}) (string, error) {
	bytes, errMarshal := bson.Marshal(entity)
	if errMarshal != nil {
		return "", errMarshal
	}
	bsonMap := make(map[string]interface{})
	errUnrmashal := bson.Unmarshal(bytes, bsonMap)
	if errUnrmashal != nil {
		return "", errUnrmashal
	}

	insertResult, err := m.collIns.InsertOne(context.TODO(), bsonMap)
	if err != nil {
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}
