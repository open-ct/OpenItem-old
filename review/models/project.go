package models

import (
	"context"
	"fmt"
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

type Project struct {
	field.DefaultField `bson:",inline"`
	Uuid               string           `json:"uuid" bson:"uuid"`
	Creator            string           `json:"creator" bson:"creator"`
	Status             int              `json:"status" bson:"status"`
	BasicInfo          ProjectBasicInfo `json:"basic_info" bson:"basic_info"`
}

// status: 0-preparing(after created), 1-processing, 2-paused, 3-finished, 4-terminated.

type ProjectBasicInfo struct {
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Requirement string   `json:"requirement" bson:"requirement"`
	Target      string   `json:"target" bson:"target"`
	GradeRange  []string `json:"grade_range" bson:"grade_range"`
	Subjects    []string `json:"subjects" bson:"subjects"`
	Summary     string   `json:"summary" bson:"summary"`
}

// Other struct for projects
type ProjectGroup struct {
	Admins     []string `json:"admin" bson:"admin"`
	Experts    []string `json:"expert" bson:"expert"`
	Assistants []string `json:"assistant" bson:"assistant"`
	Teachers   []string `json:"teachers" bson:"teachers"`
	OutExperts []string `json:"out_experts" bson:"out_experts"`
}

// ProjectMaterials: a collection about all questions / exam papers / files a project refered. just store their uuids
type ProjectMaterials struct {
	Questions  []string `json:"questions" bson:"questions"`
	ExamPapers []string `json:"exam_papers" bson:"exam_papers"`
	Files      []string `json:"files" bson:"files"`
}

// ProjectTimeTable: describe the time line of a project
type ProjectTimeTable struct {
	TimePoints []ProjectTimePoint `json:"time_points" bson:"time_points"`
}

type ProjectTimePoint struct {
	Title     string    `json:"title" bson:"title"`
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`
	Notice    string    `json:"notice" bson:"notice"` // todo: noticer config
	Comment   string    `json:"comment" bson:"comment"`
}

type ProjectFullInfo struct {
	Project
	Group     ProjectGroup     `json:"group" bson:"group"`
	TimeTable ProjectTimeTable `json:"time_table" bson:"time_table"`
	Materials ProjectMaterials `json:"materials" bson:"materials"`
	// todo:
	//Steps
	//Tasks
	//References
	//Reviews
}

/*
	init the database for project
*/
func init() {
	err := database.MgoProjects.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
			{Key: []string{"creator"}, Unique: false},
			{Key: []string{"basic_info.name"}, Unique: false},
		},
	)
	if err != nil {
		logger.Recorder.Error("[mongo project initial] create indexes error: " + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in projects collection successfully")
	return
}

func CreateEmptyProject(req *request.CreateProject) (string, int) {
	newProject := Project{
		Uuid:    utils.GenUuidV4(),
		Creator: req.UserId,
		Status:  0,
		BasicInfo: ProjectBasicInfo{
			Name:        req.Name,
			Description: req.Description,
			Requirement: req.Requirement,
			Target:      req.Target,
			GradeRange:  req.GradeRange,
			Subjects:    req.Subjects,
			Summary:     req.Summary,
		},
	}
	result, err := database.MgoProjects.InsertOne(context.Background(), &newProject)
	if err != nil {
		logger.Recorder.Error("[Mongo] create new project: " + err.Error())
		return "", response.ProjectCreateFail
	}
	// create the assignment: creator is an admin of the project
	assign := Assignment{
		DefaultField: field.DefaultField{},
		Uuid:         utils.GenUuidV4(),
		UserId:       req.UserId,
		ProjectId:    newProject.Uuid,
		Role:         1,
		Operator:     "system",
		IsConfirmed:  true,
		Status:       0,
	}
	if _, err := database.MgoAssignments.InsertOne(context.Background(), &assign); err != nil {
		logger.Recorder.Error("[Mongo] Create project's assignment error: " + err.Error())
		// delete project
		database.MgoProjects.Remove(context.Background(), bson.M{
			"uuid": newProject.Uuid,
		})
		return "", response.ProjectCreateFail
	}
	logger.Recorder.Info(fmt.Sprintf("[Project] New project and assignment has been created pid:%s ", result.InsertedID))
	return newProject.Uuid, response.SUCCESS
}

func CreateTemplateProject(req *request.CreateProject) (string, int) {
	// create an empty project

	// create 7 standard steps

	return "", response.SUCCESS
}

func UpdateProjectInfo(req *request.UpdateProjectInfo) int {
	update := bson.M{}
	if req.Name != "" {
		update["basic_info.name"] = req.Name
	}
	if req.Description != "" {
		update["basic_info.description"] = req.Description
	}
	if req.Requirement != "" {
		update["basic_info.requirement"] = req.Requirement
	}
	if req.Target != "" {
		update["basic_info.target"] = req.Target
	}
	if req.Summary != "" {
		update["basic_info.summary"] = req.Summary
	}
	if len(req.Subjects) != 0 {
		update["basic_info.subjects"] = req.Subjects
	}
	if len(req.GradeRange) != 0 {
		update["basic_info.grade_range"] = req.GradeRange
	}
	err := database.MgoProjects.UpdateOne(
		context.Background(),
		bson.M{
			"uuid": req.ProjectId,
		},
		bson.M{
			"$set": update,
		},
	)
	if err != nil {
		logger.Recorder.Error("[mongo project] update project info: " + err.Error())
		return response.ProjectUpdateFail
	}
	return response.SUCCESS
}

func GetProjectBasicInfo(pid string) (*Project, int) {
	var p Project
	err := database.MgoProjects.Find(context.Background(), bson.M{
		"uuid": pid,
	}).One(&p)
	if err != nil {
		logger.Recorder.Error("[mongo project] get project info: " + err.Error())
		return nil, response.ProjectGetInfoFail
	}
	return &p, response.SUCCESS
}

func GetProjectFullInfo(pid string) {

}
