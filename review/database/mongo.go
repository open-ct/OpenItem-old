package database

import (
	"context"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"review/logger"
)

// MongoCollections mongoDB collections name struct:
type MongoCollections struct {
	Users       string
	Assignments string
	Projects    string
	Steps       string
	Tasks       string
	FileRecords string
	Submits     string
	Audits      string
}

// MongoClient qmgo 驱动配置
var MongoClient *qmgo.Client
var (
	MgoUsers       *qmgo.Collection
	MgoProjects    *qmgo.Collection
	MgoSteps       *qmgo.Collection
	MgoTasks       *qmgo.Collection
	MgoAssignments *qmgo.Collection
	MgoFileRecords *qmgo.Collection
	MgoSubmits     *qmgo.Collection
	MgoAudits      *qmgo.Collection
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
	MgoUsers = qmgoClient.Database(mongoDbName).Collection(mongoColls.Users)
	MgoProjects = qmgoClient.Database(mongoDbName).Collection(mongoColls.Projects)
	MgoSteps = qmgoClient.Database(mongoDbName).Collection(mongoColls.Steps)
	MgoTasks = qmgoClient.Database(mongoDbName).Collection(mongoColls.Tasks)
	MgoAssignments = qmgoClient.Database(mongoDbName).Collection(mongoColls.Assignments)
	MgoFileRecords = qmgoClient.Database(mongoDbName).Collection(mongoColls.FileRecords)
	MgoSubmits = qmgoClient.Database(mongoDbName).Collection(mongoColls.Submits)
	MgoAudits = qmgoClient.Database(mongoDbName).Collection(mongoColls.Audits)

	//defer func() {
	//	if err = MongoClient.Close(context.Background()); err != nil {
	//		panic(err)
	//	}
	//}()
	return
}

// loadMongoCollectionsName load the collections name of mongo
func loadMongoCollectionsName() (*MongoCollections, error) {
	usersColl, err := beego.AppConfig.String("mongo-collections::users")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	assignments, err := beego.AppConfig.String("mongo-collections::assignments")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	proj, err := beego.AppConfig.String("mongo-collections::projects")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	steps, err := beego.AppConfig.String("mongo-collections::steps")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	tasks, err := beego.AppConfig.String("mongo-collections::tasks")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	fileRec, err := beego.AppConfig.String("mongo-collections::fileRecords")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	//ref, err := beego.AppConfig.String("mongo-collections::references")
	//if err != nil {
	//	logger.Recorder.Error("[Mongo Config] " + err.Error())
	//	return nil, err
	//}
	sub, err := beego.AppConfig.String("mongo-collections::submits")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	audit, err := beego.AppConfig.String("mongo-collections::audits")
	if err != nil {
		logger.Recorder.Error("[Mongo Config] " + err.Error())
		return nil, err
	}
	colls := &MongoCollections{
		Users:       usersColl,
		Assignments: assignments,
		Projects:    proj,
		Steps:       steps,
		Tasks:       tasks,
		FileRecords: fileRec,
		Submits:     sub,
		Audits:      audit,
	}
	return colls, nil
}

// mongoClose close the mongo connection.
func mongoClose() {
	if err := MongoClient.Close(context.Background()); err != nil {
		logger.Recorder.Error("[Mongo] Close Mongo Connection Error: " + err.Error())
	}
}

// GenMongoFilter todo: 用于生成 更新/查找 使用的filter
func GenMongoFilter(obj map[string]interface{}) bson.D {
	return bson.D{
		{"is_deleted", 0},
	}
}

/*
	下面为测试方法, 不在项目中使用
*/

// ConvertToMap 类型转换
func ConvertToMap(rawData interface{}) (map[string]interface{}, error) {
	tempBytes, err := bson.MarshalExtJSON(rawData, false, true)
	if err != nil {
		return nil, err
	}
	var jsonMap map[string]interface{}
	err = json.Unmarshal(tempBytes, &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}
