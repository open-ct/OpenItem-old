package models

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo/field"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"proj-review/constant"
	"proj-review/database"
	"proj-review/log"
	"proj-review/request"
	"proj-review/response"
	"proj-review/utils"
	"time"
)

type Project struct {
	field.DefaultField `bson:",inline"`
	Uuid               string             `json:"uuid" bson:"uuid"`
	Creator            string             `json:"creator" bson:"creator"`
	Status             int                `json:"status" bson:"status"`
	BasicInfo          ProjectBasicInfo   `json:"basic_info" bson:"basic_info"`
	Group              ProjectGroup       `json:"group" bson:"group"`
	Timetable          []ProjectTimePoint `json:"timetable" bson:"timetable"`
	Steps              []Step             `json:"steps" bson:"steps"`
}

// 项目团队
type ProjectGroup struct {
	Admins     []string `json:"admin" bson:"admin"`
	Experts    []string `json:"expert" bson:"expert"`
	Assistants []string `json:"assistant" bson:"assistant"`
	Teachers   []string `json:"teachers" bson:"teachers"`
	OutExperts []string `json:"out_experts" bson:"out_experts"`
}

// 项目基本信息
type ProjectBasicInfo struct {
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Requirement string   `json:"requirement" bson:"requirement"`
	Target      string   `json:"target" bson:"target"`
	GradeRange  []string `json:"grade_range" bson:"grade_range"`
	Subjects    []string `json:"subjects" bson:"subjects"`
	Summary     string   `json:"summary" bson:"summary"`
}

// 项目时间表(时间点数组)
type ProjectTimePoint struct {
	Name        string    `json:"name" bson:"name"`
	Time        time.Time `json:"time" bson:"time"`
	Description string    `json:"description" bson:"description"`
}

// 材料仓库
type ProjectMaterials struct {
	Questions  []string `json:"questions" bson:"questions"`
	TestPapers []string `json:"test_papers" bson:"test_papers"`
	Files      []string `json:"files" bson:"files"`
}

// todo: 其他项目项目属性

// 项目数据库操作
func init() {
	// create index
	//err := database.MgoProjects.DropIndex(context.Background(), []string{"creator"})
	//if err != nil {
	//	log.Logger.Warn("[Mongo Project] clear mongo index in project collection error: " + err.Error())
	//}
	err := database.MgoProjects.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
			{Key: []string{"creator"}, Unique: false},
			{Key: []string{"basic_info.name"}, Unique: false},
		},
	)
	if err != nil {
		log.Logger.Error("[Mongo]" + err.Error())
		return
	}
	log.Logger.Info("[Mongo] Create the index in projects collection successfully")
	return
}

// DoCreateNewProject
func DoCreateNewProject(createProjReq *request.CreateProject) (*response.ProjectDefault, bool) {
	// 只创建基本信息, 后续材料以及时间安排单独API设定
	newProj := Project{
		Uuid: utils.GenUuidV4(),
		BasicInfo: ProjectBasicInfo{
			Name:        createProjReq.Name,
			Description: createProjReq.Description,
			Requirement: createProjReq.Requirement,
			Target:      createProjReq.Target,
			GradeRange:  createProjReq.GradeRange,
			Subjects:    createProjReq.Subjects,
			Summary:     createProjReq.Summary,
		},
		Creator: createProjReq.Creator,
		Group: ProjectGroup{
			Admins: []string{createProjReq.Creator},
		},
		Status: 1, // need to define the status code. 1: project is being prepared...
	}

	// unique ?
	// may be use Transactions to create project and assignments..
	createRes, err := database.MgoProjects.InsertOne(context.Background(), &newProj)
	if err != nil {
		log.Logger.Error("[Mongo] create new project: " + err.Error())
		return &response.ProjectDefault{
			ProjectName: createProjReq.Name,
			Description: "create fail...", // todo
		}, false
	}
	if err := createOneAssignment(newProj.Creator, newProj.Uuid, 0, "system", true); err != nil {
		log.Logger.Error("[Mongo] Create project's assignment error: " + err.Error())
		// delete project
		database.MgoProjects.Remove(context.Background(), bson.M{
			"uuid": newProj.Uuid,
		})
		return &response.ProjectDefault{
			ProjectName: createProjReq.Name,
			Description: "create fail...", // todo
		}, false
	}
	log.Logger.Info(fmt.Sprintf("[Project] New project and assignment has been created %s", createRes.InsertedID))
	return &response.ProjectDefault{
		ProjectID:   newProj.Uuid,
		ProjectName: createProjReq.Name,
		Description: constant.RegisterMsg.Ok,
	}, true
}

// DoCreateTemplateProject
func DoCreateTemplateProject() {

}

// DoSearchProject
func DoSearchProject() {

}

// DoGetProject
func DoGetProject(projId string) (*response.ProjectInfo, bool) {
	proj, err := getProjectById(projId)
	if err != nil {
		log.Logger.Warn("[Project Mongo] no record: " + err.Error())
		if err == mongo.ErrNoDocuments {
			return &response.ProjectInfo{
				ProjectId:   projId,
				Description: "cannot find project.",
			}, false
		} else {
			return &response.ProjectInfo{
				ProjectId:   projId,
				Description: "something error.",
			}, false
		}
	}
	projInfo := response.ProjectInfo{
		ProjectId:          proj.Uuid,
		ProjectName:        proj.BasicInfo.Name,
		ProjectDescription: proj.BasicInfo.Description,
		Requirement:        proj.BasicInfo.Requirement,
		Target:             proj.BasicInfo.Target,
		GradeRange:         proj.BasicInfo.GradeRange,
		Subjects:           proj.BasicInfo.Subjects,
		Summary:            proj.BasicInfo.Summary,
		Creator:            proj.Creator,
		Status:             proj.Status,
		Group:              proj.Group,
		Timetable:          proj.Timetable,
		Steps:              proj.Steps,
		Description:        "ok",
		CreateAt:           proj.CreateAt,
		UpdatedAt:          proj.UpdateAt,
	}
	return &projInfo, true
}

/*
	private function
*/

func getProjectById(projId string) (Project, error) {
	var proj Project
	err := database.MgoProjects.Find(
		context.Background(),
		bson.M{
			"uuid": projId,
		},
	).One(&proj)
	if err != nil {
		return Project{}, err
	}
	return proj, nil
}
