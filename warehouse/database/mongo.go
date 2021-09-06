package database

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"warehouse/logger"
)

// In this part of PQBS, using official mongo driver to connect the mongoDB

type MongoDB struct {
	Client       *mongo.Client
	DatabaseName string
}

type MongoConfig struct {
	Host         string
	Port         int
	DatabaseName string
	User         string
	Password     string
}

var (
	TempQuestion   *mongo.Collection
	FinalQuestion  *mongo.Collection
	TempTestPaper  *mongo.Collection
	FinalTestPaper *mongo.Collection
)

var Mongo MongoDB

func init() {
	// 基本配置加载
	mongoAddress, err := beego.AppConfig.String("mongoaddr")
	if err != nil {
		fmt.Println("Load mongo config error:", err.Error())
	}
	mongoPort, err := beego.AppConfig.String("mongoport")
	if err != nil {
		fmt.Println("Load mongo config error:", err.Error())
	}
	mongoDbName, err := beego.AppConfig.String("mongodb")
	if err != nil {
		fmt.Println("Load mongo config error:", err.Error())
	}
	mongoUser, err := beego.AppConfig.String("mongouser")
	if err != nil {
		fmt.Println("Load mongo config error:", err.Error())
	}
	mongoPassword, err := beego.AppConfig.String("mongopwd")
	if err != nil {
		fmt.Println("Load mongo config error:", err.Error())
	}

	conn := fmt.Sprintf("mongodb://%s:%s", mongoAddress, mongoPort)
	clinetOptions := options.Client().ApplyURI(conn).SetAuth(options.Credential{
		AuthSource: mongoDbName,
		Username:   mongoUser,
		Password:   mongoPassword,
	})
	mongoClient, err := mongo.Connect(context.TODO(), clinetOptions)
	if err != nil {
		logger.Recorder.Error("Connect MongoDB Error" + err.Error())
		return
	}
	Mongo.Client = mongoClient
	Mongo.DatabaseName = mongoDbName

	err = loadMongoCollectionsName()
	if err != nil {
		logger.Recorder.Error("Init MongoDB Collections Error" + err.Error())
		return
	}

	logger.Recorder.Info("Connect MongoDB Successfully")
	return
}

func CloseMongo() {
	err := Mongo.Client.Disconnect(context.TODO())
	if err != nil {
		logger.Recorder.Error("Disconnect MongoDB Error")
		return
	}
	logger.Recorder.Info("Disconnect MongoDB Successfully")
	return
}

func loadMongoCollectionsName() error {
	tempQ, err := beego.AppConfig.String("mongo-collections::tempQuestions")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return err
	}
	finalQ, err := beego.AppConfig.String("mongo-collections::finalQuestions")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return err
	}
	tempTP, err := beego.AppConfig.String("mongo-collections::tempTestPapers")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return err
	}
	finalTP, err := beego.AppConfig.String("mongo-collections::finalTestPapers")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return err
	}
	TempQuestion = Mongo.Client.Database(Mongo.DatabaseName).Collection(tempQ)
	FinalQuestion = Mongo.Client.Database(Mongo.DatabaseName).Collection(finalQ)
	TempTestPaper = Mongo.Client.Database(Mongo.DatabaseName).Collection(tempTP)
	FinalTestPaper = Mongo.Client.Database(Mongo.DatabaseName).Collection(finalTP)

	return nil
}

// test database:
type Demo struct {
	Name  string `bson:"name"`
	Count int    `bson:"count"`
}

func InsertDemoMongo() {
	coll := Mongo.Client.Database(Mongo.DatabaseName).Collection("test")
	demo := Demo{
		Name:  "Hello",
		Count: 12,
	}
	insert, err := coll.InsertOne(context.TODO(), demo)
	if err != nil {
		fmt.Println("fail")
		return
	}
	fmt.Println("success", insert.InsertedID)
	return

}
