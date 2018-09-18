package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
)

// Weather : structure for mongo record
type Weather struct {
	SerialNo  string  `bson:"serialNo" json:"serialNo,omitempty"`
	TimeStamp string  `bson:"timeStamp" json:"timeStamp,omitempty"`
	Temp      float32 `bson:"temparature" json:"temparature,omitempty"`
	Humid     float32 `bson:"humidity" json:"humidity,omitempty"`
	PM2       float32 `bson:"pm2" json:"pm2,omitempty"`
	Hchco     float32 `bson:"hchco" json:"hchco,omitempty"`
	Ozone     float32 `bson:"ozone" json:"ozone,omitempty"`
	Co2       float32 `bson:"co2" json:"co2,omitempty"`
	Tvoc      float32 `bson:"tvoc" json:"tvoc,omitempty"`
}

var mongoClient *mongo.Client
var err error
var database *mongo.Database
var collection *mongo.Collection
var weatherDeatils []Weather

func initMongo() bool {
	//setting connection string
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoConfig.username, mongoConfig.password,
		mongoConfig.host, mongoConfig.port)
	if mongoConfig.username == "" && mongoConfig.password == "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s",
			mongoConfig.host, mongoConfig.port)
	} else if mongoConfig.username == "" || mongoConfig.password == "" {
		fmt.Println("Please provide MONGO_USER and MONGO_PASSWORD")
		return false
	}
	fmt.Println("Connection String : " + connectionString)

	//connecting with mongo db
	mongoClient, err = mongo.Connect(context.Background(), connectionString)
	if err != nil {
		fmt.Println("Mongo connection error occured!")
		fmt.Printf("Connection String : %s\n", connectionString)
		return false
	}

	//setting database and collection
	database = mongoClient.Database(mongoConfig.database)
	collection = database.Collection("sensor_data")

	return true
}

func getData(serialNo string, pageSize string, pageNum string) []Weather {
	var weatherDeatils []Weather
	intPageSize, err := strconv.Atoi(pageSize)
	intPageNum, err := strconv.Atoi(pageNum)
	skips := intPageSize * (intPageNum - 1)
	if err != nil {
		fmt.Println(err)
	}
	cursor, err := collection.Find(context.Background(), bson.NewDocument(bson.EC.String("serialNo", serialNo)), findopt.Skip(int64(skips)), findopt.Limit(int64(intPageSize)))
	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(context.Background()) {
		var result Weather
		cursor.Decode(&result)
		weatherDeatils = append(weatherDeatils, result)
	}
	fmt.Println(weatherDeatils)
	return weatherDeatils
}

func saveData(weather Weather) bool {
	fmt.Println(weather)
	result, err := collection.InsertOne(context.Background(), weather)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("Inserted : %s", result.InsertedID)
	return true

}
