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
// @Description 文件上传, 使用post form格式上传, 会自动解析token获得对应的上传者id
// @Param   token header string true "user token get at login"
// @Param   file formData file true "文件名"
// @Param   description formData string false "文件注释和说明"
// @Param   tags formData string false "文件标签(文件类型即为文件的后缀名, 自动解析)"
// @Param   source_project formData string true "上传文件对应的项目id, 查询使用"
// @Success 200 {object} response.Default
// @Failure 400 "invalid file"
// @router / [post]
func (f *FileController) UploadFile() {
	file, fileHeader, err := f.GetFile("file")
	if err != nil {
		logger.Recorder.Warning("[file] get file from post-request error: " + err.Error())
		f.respondJson(http.StatusBadRequest, response.FAIL, "get file failed")
		return
	}
	defer file.Close()
	fmt.Println(f.GetString("description"), f.GetString("tags"))
	fileDescription := f.GetString("description")
	fileTags := strings.Split(f.GetString("tags"), ",")
	fileSourceProject := f.GetString("source_project")
	uploader, err := parseUserToken(f.Ctx.Request.Header["Token"][0])
	// if no token, access module will block the request.
	uploadRequest := request.UploadFile{
		UserId:        uploader,
		FileName:      fileHeader.Filename,
		Type:          path.Ext(fileHeader.Filename),
		SourceProject: fileSourceProject,
		Description:   fileDescription,
		Tags:          fileTags,
	}
	fileRecord, code := models.CreateNewFileRecord(&uploadRequest)
	if code == response.SUCCESS {
		err := f.SaveToFile("file", fileRecord.Path)
		if err != nil {
			logger.Recorder.Error("[upload file] error:" + err.Error())
			// delete the file record
			models.DeleteFile(fileRecord.Uuid)
			f.respondJson(http.StatusOK, response.FAIL, "上传文件错误", fileRecord)
			return
		}
	}
	f.respondJson(http.StatusOK, code, "", fileRecord)
	return
}

// @Title DownloadFile
// @Description download a file by file id(根据文件id下载文件)
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
// @Description 只获取文件信息, 不执行下载
// @Param   token header string true "user token get at login"
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
// @Description 条件搜素文件, 搜索结果 (待完善)
// @Param   token header string true "user token get at login"
// @Param   json body request.SearchFile true "搜索条件"
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
// @Description 删除文件上传记录(不在磁盘存储中删除文件)
// @Param   token header string true "user token get at login"
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
