package models

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"proj-review/constant"
	"proj-review/database"
	"proj-review/log"
	"proj-review/request"
	"proj-review/response"
	"proj-review/utils"
	"time"
)

// FileItem
type FileItem struct {
	Uuid        string    `json:"uuid" bson:"uuid"`
	Name        string    `json:"name" bson:"name"`
	Type        string    `json:"type" bson:"type"`
	Description string    `json:"description" bson:"description"`
	Tags        []string  `json:"tags" bson:"tags"`
	Path        string    `json:"path" bson:"path"`
	Owner       string    `json:"owner" bson:"owner"` // uploader's uuid
	CreatedTime time.Time `json:"created_time" bson:"created_time"`
	IsDeleted   bool      `json:"is_deleted" bson:"is_deleted"`
}

// FileStoreConfig define the location of files to save.
type FileStoreConfig struct {
	RootPath   string
	NameFormat string
}

var FileStore FileStoreConfig

func init() {
	// 加载文件存储目录
	fileStoreRootPath, err := beego.AppConfig.String("filerootpath")
	if err != nil {
		fmt.Println("Load file store config error:", err.Error())
	}
	FileStore.RootPath = fileStoreRootPath
	// create the index of files-collections
	err = database.MgoFileRecords.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true}, // uuid is unique
			{Key: []string{"tags", "name", "type"}, Unique: false},
		},
	)
	if err != nil {
		log.Logger.Error("[Mongo]" + err.Error())
		return
	}
	log.Logger.Info("[Mongo] Create the index in file-records collection successfully")
	return
}

// DoRecordFileInfo 在mongo中记录文件索引表
func DoRecordFileInfo(uploadReq *request.UploadFile) (*response.FileDefault, bool) {
	fileID := utils.GenUuidV4()
	file := FileItem{
		Uuid:        fileID,
		Name:        uploadReq.FileName,
		Type:        uploadReq.Type,
		Description: uploadReq.Description,
		Tags:        uploadReq.Tags,
		Path:        genFilesPath(fileID, uploadReq.Type),
		Owner:       uploadReq.UserId, // record the uploader's id
		CreatedTime: time.Now(),
	}
	insertRes, err := database.MgoFileRecords.InsertOne(context.Background(), file)
	if err != nil {
		log.Logger.Warn("[File Insert] " + err.Error())
		return &response.FileDefault{
			FileID:      "",
			FileName:    uploadReq.FileName,
			Description: constant.FileMsg.Fail,
		}, false
	}
	log.Logger.Info("[Mongo Insert] " + fmt.Sprintf("%s", insertRes.InsertedID))
	return &response.FileDefault{
		FileID:      file.Uuid,
		FileName:    uploadReq.FileName,
		FileType:    uploadReq.Type,
		FilePath:    file.Path,
		Description: constant.FileMsg.Ok,
	}, true
}

// DoGetFileInfo 获取文件记录信息
func DoGetFileInfo(fileUuid string) (*response.FindFile, bool) {
	var fileInfo FileItem
	err := database.MgoFileRecords.Find(context.Background(), bson.M{"uuid": fileUuid}).One(&fileInfo)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return &response.FindFile{
			FileID:      fileUuid,
			Description: constant.FileMsg.Fail,
		}, false
	}
	return &response.FindFile{
		FileID:          fileInfo.Uuid,
		FileName:        fileInfo.Name,
		FileType:        fileInfo.Type,
		FileTags:        fileInfo.Tags,
		FilePath:        genFilesPath(fileUuid, fileInfo.Type),
		FileDescription: fileInfo.Description,
		Description:     constant.FileMsg.Ok,
	}, true
}

// todo: 文件搜素
func DoSearchFile(searchReq *request.SearchFile) (*response.SearchFiles, bool) {
	// operations
	return nil, true
}

/*
	additional functions:
*/
// genFilesPath generate the files saving-path and saving-filename.
func genFilesPath(fileID string, fileType string) string {
	t := time.Now()
	dateString := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
	todayPath := FileStore.RootPath + dateString
	// 如果没有目录, 需要创建
	if !utils.IsDirExists(todayPath) {
		fmt.Println("dir not exist")
		utils.CreateDateDir(FileStore.RootPath)
	}
	return todayPath + "/" + fileID + fileType
}
