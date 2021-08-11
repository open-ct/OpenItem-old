package controllers

import (
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"path"
	"proj-review/constant"
	"proj-review/log"
	"proj-review/models"
	"proj-review/request"
	"proj-review/response"
)

type FileController struct {
	beego.Controller
}

// respondJson make respond (in json typr)
func (f *FileController) respondJson(httpCode int, opCode int, message string, data ...interface{}) {
	f.Ctx.Output.SetStatus(httpCode)
	var d interface{}
	if len(data) == 1 {
		d = data[0]
	} else {
		d = data
	}
	resp := response.GenResponse(
		opCode,
		message,
		d,
	)
	f.Data["json"] = resp
	f.ServeJSON()
}

// UploadFile (need to record the user id)
func (f *FileController) UploadFile() {
	// 获取前端上传文件
	file, fileHeader, err := f.GetFile("file_name")
	if err != nil || fileHeader == nil {
		log.Logger.Warn("[File] " + err.Error())
		f.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.UploadFileFail,
		)
		return
	}
	// confirm the user id
	userId, err := parseUserToken(f.GetString("token"))
	if err != nil {
		log.Logger.Warn("[File] " + err.Error())
		f.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.UploadFileFail, // unknown upload user ... operation denied
		)
	}
	fileReq := request.UploadFile{
		UserId:      userId,
		FileName:    fileHeader.Filename,
		Type:        path.Ext(fileHeader.Filename),
		Description: "",
		Tags:        []string{},
	}
	defer file.Close()
	// 数据库中记录对应文件上传信息
	uploadResp, ok := models.DoRecordFileInfo(&fileReq)
	if !ok {
		f.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.UploadFileFail,
			uploadResp,
		)
	} else {
		f.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			constant.BasicMsg.UploadFileSuccess,
			uploadResp,
		)
	}
	f.SaveToFile("file_name", uploadResp.FilePath)
	return
}

// DownloadFile 根据Uuid下载文件
func (f *FileController) DownloadFile() {
	fileId := f.GetString("file_id")
	if fileId == "" {
		log.Logger.Warn("[File] " + errors.New("Invalid file ID").Error())
		f.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.DownloadFileFail,
		)
		return
	}
	// 获取下载路径
	fileInfoResp, ok := models.DoGetFileInfo(fileId)
	if !ok {
		f.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.GetFileInfoFail,
			fileInfoResp,
		)
	} else {
		// 开始下载
		f.Ctx.Output.Download(fileInfoResp.FilePath, fileInfoResp.FileName)
	}
	return
}

// GetFileInfo 根据ID获取文件信息
func (f *FileController) GetFileInfo() {
	fileId := f.GetString("file_id")
	if fileId == "" {
		log.Logger.Warn("[File] " + errors.New("Invalid file ID").Error())
		f.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.GetInfoFail,
		)
		return
	}
	fileInfoResp, ok := models.DoGetFileInfo(fileId)
	if !ok {
		f.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.GetFileInfoFail,
			fileInfoResp,
		)
	} else {
		f.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			constant.BasicMsg.GetFileInfoSuccess,
			fileInfoResp,
		)
	}
	return
}

// todo: 文件搜素
func (f *FileController) SearchFile() {

}
