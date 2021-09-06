package models

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/qiniu/qmgo/field"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"review/database"
	"review/logger"
	"review/request"
	"review/response"
	"review/utils"
	"time"
)

// FileItem: define the file-record in mongodb
type FileItem struct {
	// qmgo default filed (id, create, update)
	field.DefaultField `bson:",inline"`
	Uuid               string   `json:"uuid" bson:"uuid"`
	Name               string   `json:"name" bson:"name"`
	Type               string   `json:"type" bson:"type"`
	Base               string   `json:"base" bson:"base"`
	IsPublic           bool     `json:"is_public" bson:"is_public"`
	Belongs            []string `json:"belongs" bson:"belongs"`
	Description        string   `json:"description" bson:"description"`
	Tags               []string `json:"tags" bson:"tags"`
	Path               string   `json:"path" bson:"path"`
	Owner              string   `json:"owner" bson:"owner"` // uploader's uuid
}

// fileStoreConfig: define the location of files to save.
type fileStoreConfig struct {
	RootPath   string
	NameFormat string
}

var fileStore fileStoreConfig

func init() {
	// 加载文件存储目录
	fileStoreRootPath, err := beego.AppConfig.String("filerootpath")
	if err != nil {
		fmt.Println("Load file store config error:", err.Error())
	}
	fileStore.RootPath = fileStoreRootPath
	// create the index of files-collections
	err = database.MgoFileRecords.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
			{Key: []string{"tags"}, Unique: false},
			{Key: []string{"name"}, Unique: false},
			{Key: []string{"type"}, Unique: false},
		},
	)
	if err != nil {
		logger.Recorder.Error("[Mongo]" + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in file-records collection successfully")
	return
}

// CreateNewFileRecord: create an index of file in mongodb, return: record-id, op-code
func CreateNewFileRecord(req *request.UploadFile) (*FileItem, int) {
	fileUuid := utils.GenUuidV4()
	file := FileItem{
		Uuid:        fileUuid,
		Name:        req.FileName,
		Type:        req.Type,
		Base:        "root",
		IsPublic: req.IsPublic,
		Belongs: req.Belongs,
		Description: req.Description,
		Tags:        req.Tags,
		Path:        genFilesPath(fileUuid, req.Type),
		Owner:       req.UserId, // record the uploader's id
	}
	insertRes, err := database.MgoFileRecords.InsertOne(context.Background(), &file)
	if err != nil {
		logger.Recorder.Warn("[File upload (mongo insert new file record failed)] " + err.Error())
		return nil, response.DatabaseInsertError
	}
	logger.Recorder.Info("[Mongo Insert] " + fmt.Sprintf("%s", insertRes.InsertedID))
	return &file, response.SUCCESS
}

// GetFileInfo: get file information by file-uuid
func GetFileInfo(fileUuid string) (*FileItem, int) {
	var fileInfo FileItem
	err := database.MgoFileRecords.Find(context.Background(), bson.M{"uuid": fileUuid}).One(&fileInfo)
	if err != nil {
		logger.Recorder.Warn("err: ", err.Error())
		return nil, response.DatabaseNoRecord
	}
	return &fileInfo, response.SUCCESS
}

// SearchFiles
func SearchFiles(searchReq *request.SearchFile) (*[]FileItem, int) {
	// operations
	var files []FileItem
	filter := []bson.M{}

	if searchReq.FileName != "" {
		filter = append(filter, bson.M{
			"file_name": searchReq.Type,
		})
	}
	if searchReq.Type != "" {
		filter = append(filter, bson.M{
			"type": searchReq.Type,
		})
	}
	searchFilter := bson.M{
		"$or":filter,
	}
	err := database.MgoFileRecords.Find(context.Background(), searchFilter).All(&files)
	if err != nil {
		logger.Recorder.Error("[Mongo Search File] " + err.Error())
		return nil, response.DatabaseNoRecord
	}

	return &files, response.SUCCESS
}

// DeleteFile: delete the file-record in mongodb (keep on disk currently)
func DeleteFile(fileID string) int {
	// todo: delete file on server disk.
	err := database.MgoFileRecords.Remove(context.Background(), bson.M{
		"uuid": fileID,
	})
	if err != nil {
		logger.Recorder.Error("[File Delete] " + err.Error())
		return response.DatabaseDeleteError
	}
	return response.SUCCESS
}

/*
	additional functions:
*/
// genFilesPath generate the files saving-path and saving-filename.
func genFilesPath(fileID string, fileType string) string {
	t := time.Now()
	dateString := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
	todayPath := fileStore.RootPath + dateString
	// 如果没有目录, 需要创建
	if !utils.IsDirExists(todayPath) {
		fmt.Println("dir not exist")
		utils.CreateDateDir(fileStore.RootPath)
	}
	return todayPath + "/" + fileID + fileType
}

// insertDemoFile (test)
func insertDemoFile() {
	demoName := "robot-demo"
	fileName := utils.GenSha256(demoName)
	count, err := database.MgoFileRecords.Find(context.Background(), bson.M{
		"name": fileName,
	}).Count()
	if err != nil || count == 0 {
		toIndsert := FileItem{
			Uuid:        utils.GenUuidV4(),
			Name:        fileName,
			Type:        "empty",
			Description: "this is a demo file, develop testing",
			Path:        "no path",
		}
		result, err := database.MgoFileRecords.InsertOne(context.Background(), &toIndsert)
		if err != nil {
			logger.Recorder.Error("[Mongo File] Init error: " + err.Error())
		}
		logger.Recorder.Info(fmt.Sprintf("[Mongo File] init insert ok: %s", result.InsertedID))
		return
	}
	logger.Recorder.Info("[Mongo File] No Insert operation")
	return

}
