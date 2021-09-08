package database

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/qiniu/qmgo"
	"warehouse/logger"
)

// MongoCollections mongoDB collections name struct:
type MongoCollections struct {
	TempQuestions  string
	FinalQuestions string
	TempTestPapers string
	FinalTestPaper string
}

// MongoClient qmgo 驱动配置
var MongoClient *qmgo.Client
var (
	MgoTempQuestions  *qmgo.Collection
	MgoFinalQuestions *qmgo.Collection
	MgoTempTestPaper  *qmgo.Collection
	MgoFinalTestPaper *qmgo.Collection
)

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

	// 配置 collections name, 若配置失败, 不连接 mongo 直接返回
	mongoColls, err := loadMongoCollectionsName()
	if err != nil {
		fmt.Println("Load mongo collections config error: ", err.Error())
		return
	}

	// 配置完成后再开启mongo 连接
	conn := fmt.Sprintf("mongodb://%s:%s", mongoAddress, mongoPort)

	// 使用 qmgo 连接
	qmgoClient, err := qmgo.NewClient(
		context.Background(),
		&qmgo.Config{
			Uri:      conn,
			Database: mongoDbName,
			Auth: &qmgo.Credential{
				AuthSource: mongoDbName,
				Username:   mongoUser,
				Password:   mongoPassword,
			},
		},
	)
	MongoClient = qmgoClient

	if err != nil {
		logger.Recorder.Error("[Mongo-qmgo]" + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo-qmgo] Connected successfully")

	// 对接collections
	MgoTempQuestions = qmgoClient.Database(mongoDbName).Collection(mongoColls.TempQuestions)
	MgoFinalQuestions = qmgoClient.Database(mongoDbName).Collection(mongoColls.FinalQuestions)
	MgoTempTestPaper = qmgoClient.Database(mongoDbName).Collection(mongoColls.TempTestPapers)
	MgoFinalQuestions = qmgoClient.Database(mongoDbName).Collection(mongoColls.FinalTestPaper)

	return
}

// loadMongoCollectionsName load the collections name of mongo
func loadMongoCollectionsName() (*MongoCollections, error) {
	tempQ, err := beego.AppConfig.String("mongo-collections::tempQuestions")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	finalQ, err := beego.AppConfig.String("mongo-collections::finalQuestions")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	tempTP, err := beego.AppConfig.String("mongo-collections::tempTestPapers")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	finalTP, err := beego.AppConfig.String("mongo-collections::finalTestPapers")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	colls := &MongoCollections{
		TempQuestions:  tempQ,
		FinalQuestions: finalQ,
		TempTestPapers: tempTP,
		FinalTestPaper: finalTP,
	}
	return colls, nil
}

// mongoClose close the mongo connection.
func mongoClose() {
	if err := MongoClient.Close(context.Background()); err != nil {
		logger.Recorder.Error("[Mongo] Close Mongo Connection Error: " + err.Error())
	}
}
