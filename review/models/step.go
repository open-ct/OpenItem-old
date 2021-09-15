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
)

type Step struct {
	field.DefaultField `bson:",inline"`
	Uuid               string             `json:"uuid" bson:"uuid"`
	ProjectId          string             `json:"project_id" bson:"project_id"`
	Index              int                `json:"index" bson:"index"`
	Name               string             `json:"name" bson:"name"`
	Description        string             `json:"description" bson:"description"`
	Requirement        string             `json:"requirement" bson:"requirement"`
	Status             int                `json:"status" bson:"status"`
	Deadline           int64              `json:"deadline" bson:"deadline"`
	Timetable          []ProjectTimePoint `json:"timetable" bson:"timetable"`
	Creator            string             `json:"creator" bson:"creator"`
	Attachments        []string           `json:"attachments" bson:"attachments"` // uuid of files
}

// Status: 0-not start, 1-processing, 2-paused, 3-finished

func init() {
	// clean org indexes.
	database.MgoSteps.DropAllIndexes(context.Background())
	err := database.MgoSteps.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
			{Key: []string{"project_id"}, Unique: false},
			{Key: []string{"project_id", "name"}, Unique: true},
			{Key: []string{"project_id", "index"}, Unique: true},
		},
	)
	if err != nil {
		logger.Recorder.Error("[Mongo Steps] " + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in steps collection successfully")
	return
}

func CreateOneStep(req *request.CreateStep) (string, int) {
	newStep := Step{
		Uuid:        utils.GenUuidV4(),
		ProjectId:   req.ProjectId,
		Index:       req.Index,
		Name:        req.Name,
		Description: req.Description,
		Requirement: req.Requirement,
		Deadline:    req.Deadline,
		Status:      0,
		Creator:     req.Creator,
	}
	newTimeTable := []ProjectTimePoint{}
	for _, timePoint := range req.TimeTable {
		point := ProjectTimePoint{
			Title:     timePoint.Title,
			StartTime: timePoint.StartTime,
			EndTime:   timePoint.EndTime,
			Notice:    timePoint.Notice,
			Comment:   timePoint.Comment,
		}
		newTimeTable = append(newTimeTable, point)
	}
	newStep.Timetable = newTimeTable

	insert, err := database.MgoSteps.InsertOne(context.Background(), &newStep)
	if err != nil {
		logger.Recorder.Warn("[File upload (mongo create new step  failed)] " + err.Error())
		return "", response.StepCreateFail
	}
	logger.Recorder.Info("[Mongo Insert] " + fmt.Sprintf("%s", insert.InsertedID))
	return newStep.Uuid, response.SUCCESS
}

func GetStepInfo(sid string) (*Step, int) {
	var step Step
	err := database.MgoSteps.Find(context.Background(), bson.M{
		"uuid": sid,
	}).One(&step)
	if err != nil {
		logger.Recorder.Warn("err: ", err.Error())
		return nil, response.StepGetInfoFail
	}
	return &step, response.SUCCESS
}

func GetAllStepsInProject(pid string) (*[]Step, int) {
	var steps []Step
	err := database.MgoSteps.Find(context.Background(), bson.M{
		"project_id": pid,
	}).All(&steps)
	if err != nil {
		logger.Recorder.Warn("err: " + err.Error())
		return nil, response.StepGetInfoFail
	}
	return &steps, response.SUCCESS
}

func UploadStepAttachments(req *request.AddStepAttachment) int {
	err := database.MgoSteps.UpdateOne(
		context.Background(),
		bson.M{
			"uuid": req.StepId,
		},
		bson.M{
			"$set": bson.M{
				"attachments": req.FilesIds,
			},
		},
	)
	if err != nil {
		logger.Recorder.Warning("[mongo] add attachments for step: " + err.Error())
		return response.StepAddAttachmentsFail
	}
	return response.SUCCESS
}

func UpdateStepInfo(req *request.UpdateStepInfo) int {
	var oldStep Step
	err := database.MgoSteps.Find(context.Background(), bson.M{
		"uuid": req.StepId,
	}).One(&oldStep)
	if err != nil {
		logger.Recorder.Warning("[mongo] update step information: " + err.Error())
		return response.StepUpdateInfoFail
	}
	filter := make(map[string]interface{})
	if req.NewName != "" {
		filter["name"] = req.NewName
	}
	if req.NewDescription != "" {
		filter["description"] = req.NewDescription
	}
	if req.NewRequirement != "" {
		filter["requirement"] = req.NewRequirement
	}
	if req.NewDeadline != oldStep.Deadline {
		filter["deadline"] = req.NewDeadline
	}
	err = database.MgoSteps.UpdateOne(context.Background(),
		bson.M{"uuid": req.StepId},
		bson.M{
			"$set": filter,
		})
	if err != nil {
		logger.Recorder.Warning("[mongo] update step information: " + err.Error())
		return response.StepUpdateInfoFail
	}
	return response.SUCCESS
}

func SetStepStatus(req *request.SetStepStatus) int {
	err := database.MgoSteps.UpdateOne(context.Background(),
		bson.M{"uuid": req.StepId},
		bson.M{
			"$set": bson.M{
				"status": req.NewStatus,
			},
		})
	if err != nil {
		logger.Recorder.Warning("[mongo] change step status: " + err.Error())
		return response.StepChangeStatusFail
	}
	return response.SUCCESS
}

func SetStepTimePoint(req *request.SetStepTimePoint) (*[]ProjectTimePoint, int) {
	// get step data:
	var step Step
	err := database.MgoSteps.Find(context.Background(), bson.M{
		"uuid": req.StepId,
	}).One(&step)
	if err != nil {
		logger.Recorder.Warning("[mongo] address the step error: " + err.Error())
		return nil, response.StepSetTimePointFail
	}
	newTimeTable := step.Timetable
	newTimePoint := ProjectTimePoint{
		Title:     req.Info.Title,
		StartTime: req.Info.StartTime,
		EndTime:   req.Info.EndTime,
		Notice:    req.Info.Notice,
		Comment:   req.Info.Comment,
	}
	if req.PointIndex < 0 || req.PointIndex >= len(step.Timetable) {
		newTimeTable = append(newTimeTable, newTimePoint)
		err := database.MgoSteps.UpdateOne(context.Background(), bson.M{"uuid": req.StepId}, bson.M{
			"$set": bson.M{
				"timetable": newTimeTable,
			},
		})
		if err != nil {
			logger.Recorder.Warning("[mongo] append step time point error: " + err.Error())
			return nil, response.StepSetTimePointFail
		}
		return &newTimeTable, response.SUCCESS
	}
	newTimeTable[req.PointIndex] = newTimePoint
	err = database.MgoSteps.UpdateOne(context.Background(),
		bson.M{"uuid": req.StepId},
		bson.M{
			"$set": bson.M{
				"timetable": newTimeTable,
			},
		})
	if err != nil {
		logger.Recorder.Warning("[mongo] append step time point error: " + err.Error())
		return nil, response.StepSetTimePointFail
	}
	return &newTimeTable, response.SUCCESS
}

func DeleteStepTimePoint(req *request.DeleteStepTimePoint) int {
	var step Step
	err := database.MgoSteps.Find(context.Background(), bson.M{
		"uuid": req.StepId,
	}).One(&step)
	if err != nil {
		logger.Recorder.Warning("[mongo] address the step error: " + err.Error())
		return response.StepDeleteTimePointFail
	}
	if req.PointIndex >= len(step.Timetable) || req.PointIndex < 0 {
		return response.StepDeleteTimePointFail
	}
	newTimeTable := step.Timetable
	// delete array element
	newTimeTable = append(newTimeTable[:req.PointIndex], newTimeTable[req.PointIndex+1:]...)
	err = database.MgoSteps.UpdateOne(context.Background(),
		bson.M{"uuid": req.StepId},
		bson.M{
			"$set": bson.M{
				"timetable": newTimeTable,
			},
		})
	if err != nil {
		logger.Recorder.Warning("[mongo] delete time point error: " + err.Error())
		return response.StepDeleteTimePointFail
	}
	return response.SUCCESS
}

// todo:
func GetStepDataStatistic() {

}

func DeleteStep(stepId string) int {
	// todo:
	err := database.MgoSteps.Remove(context.Background(), bson.M{
		"uuid": stepId,
	})
	if err != nil {
		logger.Recorder.Warning("[mongo] delete step error: " + err.Error())
		return response.StepDeleteFail
	}
	return response.SUCCESS
}
