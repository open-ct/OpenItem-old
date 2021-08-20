package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"path"
	"review/logger"
	"review/models"
	"review/request"
	"review/response"
	"strings"
)

// Operations of files
type FileController struct {
	beego.Controller
}

// respondJson: do response for user operations
func (f *FileController) respondJson(httpCode int, opCode int, message string, data ...interface{}) {
	f.Ctx.Output.SetStatus(httpCode)
	var d interface{}
	if len(data) == 1 {
		d = data[0]
	} else {
		d = data
	}
	resp := response.GenResponse(opCode, message, d)
	f.Data["json"] = resp
	f.ServeJSON()
}

// @Title UploadFile
// @Description upload a file, need use token to define the creater(owner)
// @Param   token header string true "user token get at login"
// @Param   filename formData file true "file user want to upload"
// @Param   description formData string false "the details and comment of file to uploaded"
// @Param   tags formData string false "file's tags"
// @Success 200 {object} response.Default
// @Failure 400 "invalid file"
// @router / [post]
func (f *FileController) UploadFile() {
	file, fileHeader, err := f.GetFile("filename")
	if err != nil {
		logger.Recorder.Warning("[file] get file from post-request error: " + err.Error())
		f.respondJson(http.StatusBadRequest, response.FAIL, "get file failed")
		return
	}
	defer file.Close()
	fmt.Println(f.GetString("description"), f.GetString("tags"))
	fileDescription := f.GetString("description")
	fileTags := strings.Split(f.GetString("tags"), ",")
	uploader, err := parseUserToken(f.Ctx.Request.Header["Token"][0])
	// if no token, access module will block the request.
	uploadRequest := request.UploadFile{
		UserId:      uploader,
		FileName:    fileHeader.Filename,
		Type:        path.Ext(fileHeader.Filename),
		Description: fileDescription,
		Tags:        fileTags,
	}
	fileRecord, code := models.CreateNewFileRecord(&uploadRequest)
	f.respondJson(http.StatusOK, code, "", fileRecord)
	if code == response.SUCCESS {
		f.SaveToFile("filename", fileRecord.Path)
	}
	return
}

// @Title DownloadFile
// @Description download a file by file id
// @Param   token header string true "user token get at login"
// @Param   fid path string true "file uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid file uuid"
// @router /:fid [get]
func (f *FileController) DownloadFile() {
	fid := f.GetString(":fid")
	if fid == "" {
		f.respondJson(http.StatusBadRequest, response.FAIL, "invalid file id")
		return
	}
	fileInfo, code := models.GetFileInfo(fid)
	if code == response.SUCCESS {
		f.Ctx.Output.Download(fileInfo.Path, fileInfo.Name)
		return
	}
	f.respondJson(http.StatusOK, code, "", fileInfo)
	return
}

// @Title GetFileInfo
// @Description just get file information, do not download file
// @Param   token header string false "user token get at login"
// @Param   fid path string true "file uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid file uuid"
// @router /info/:fid [get]
func (f *FileController) GetFileInfo() {
	fid := f.GetString(":fid")
	if fid == "" {
		f.respondJson(http.StatusBadRequest, response.FAIL, "invalid file id")
		return
	}
	fileInfo, code := models.GetFileInfo(fid)
	f.respondJson(http.StatusOK, code, "", fileInfo)
	return
}

// @Title SearchFiles
// @Description search files using conditions
// @Param   token header string false "user token get at login"
// @Param   json body request.SearchFile true "the search conditions"
// @Success 200 {object} response.Default
// @Failure 400 "invalid conditions (parse body failed)"
// @router /search [post]
func (f *FileController) SearchFiles() {
	searchRequest := new(request.SearchFile)
	err := unmarshalBody(f.Ctx.Input.RequestBody, searchRequest)
	if err != nil {
		logger.Recorder.Warning("[file search] parse search conditions error: " + err.Error())
		f.respondJson(http.StatusBadRequest, response.FAIL, "parse search conditions error")
		return
	}
	searchResult, code := models.SearchFiles(searchRequest)
	f.respondJson(http.StatusOK, code, "", searchResult)
	return
}

// @Title DeleteFile
// @Description delete a file record in database (keep on the disk)
// @Param   token header string false "user token get at login"
// @Param   fid path string true "the file's uuid you want to delete"
// @Success 200 {object} response.Default
// @Failure 400 "invalid file id"
// @router /:fid [delete]
func (f *FileController) DeleteFile() {
	fid := f.GetString(":fid")
	if fid == "" {
		f.respondJson(http.StatusBadRequest, response.FAIL, "invalid file id")
		return
	}
	code := models.DeleteFile(fid)
	f.respondJson(http.StatusOK, code, "")
	return
}

// todo: trace files updating
