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

type Assignment struct {
	field.DefaultField `bson:",inline"`
	Uuid               string `json:"uuid" bson:"uuid"`
	UserId             string `json:"user_id" bson:"user_id"`
	ProjectId          string `json:"project_id" bson:"project_id"`
	Role               int    `json:"role" bson:"role"`
	Operator           string `json:"operator" bson:"operator"`
	IsConfirmed        bool   `json:"is_confirmed" bson:"is_confirmed"`
	Status             int    `json:"status" bson:"status"` // todo: modify it usage
}

// Role: 0-system admin, 1-project admin, 2-expert, 3-subject assistant, 4-teacher, 5-out expert

func init() {
	// create indexes in mongo collections
	err := database.MgoAssignments.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
			{Key: []string{"user_id"}, Unique: false},
			{Key: []string{"project_id"}, Unique: false},
			{Key: []string{"user_id", "project_id"}, Unique: true},
		},
	)
	if err != nil {
		logger.Recorder.Error("[Mongo Assignment] " + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in assignments collection successfully")
	return
}

// MakeOneAssignment
func MakeOneAssignment(req *request.MakeOneAssignment) (string, int) {
	newAssign := Assignment{
		DefaultField: field.DefaultField{},
		Uuid:         utils.GenUuidV4(),
		UserId:       req.UserId,
		ProjectId:    req.ProjectId,
		Role:         req.Role,
		Operator:     req.Operator,
		IsConfirmed:  false,
		Status:       0,
	}
	result, err := database.MgoAssignments.InsertOne(context.Background(), &newAssign)
	if err != nil {
		logger.Recorder.Error("[Mongo Assignment] " + err.Error())
		return "", response.AssignmentCreateFail
	}
	logger.Recorder.Info(fmt.Sprintf("Created new assignment: %s", result.InsertedID))
	return newAssign.Uuid, response.SUCCESS
}

// MakeAssignments
func MakeAssignments(req *request.MakeAssignmentGroup) int {
	// use transaction to create a group
	ctx := context.Background()
	callback := func(sessCtx context.Context) (interface{}, error) {
		for _, item := range req.Admins {
			newAssign := Assignment{
				DefaultField: field.DefaultField{},
				Uuid:         utils.GenUuidV4(),
				UserId:       item,
				ProjectId:    req.ProjectId,
				Role:         1,
				Operator:     req.Operator,
				IsConfirmed:  false,
				Status:       0,
			}
			_, err := database.MgoAssignments.InsertOne(context.Background(), &newAssign)
			if err != nil {
				return nil, err
			}
		}
		for _, item := range req.Experts {
			newAssign := Assignment{
				DefaultField: field.DefaultField{},
				Uuid:         utils.GenUuidV4(),
				UserId:       item,
				ProjectId:    req.ProjectId,
				Role:         2,
				Operator:     req.Operator,
				IsConfirmed:  false,
				Status:       0,
			}
			_, err := database.MgoAssignments.InsertOne(context.Background(), &newAssign)
			if err != nil {
				return nil, err
			}
		}
		for _, item := range req.Assistants {
			newAssign := Assignment{
				DefaultField: field.DefaultField{},
				Uuid:         utils.GenUuidV4(),
				UserId:       item,
				ProjectId:    req.ProjectId,
				Role:         3,
				Operator:     req.Operator,
				IsConfirmed:  false,
				Status:       0,
			}
			_, err := database.MgoAssignments.InsertOne(context.Background(), &newAssign)
			if err != nil {
				return nil, err
			}
		}
		for _, item := range req.Teachers {
			newAssign := Assignment{
				DefaultField: field.DefaultField{},
				Uuid:         utils.GenUuidV4(),
				UserId:       item,
				ProjectId:    req.ProjectId,
				Role:         4,
				Operator:     req.Operator,
				IsConfirmed:  false,
				Status:       0,
			}
			_, err := database.MgoAssignments.InsertOne(context.Background(), &newAssign)
			if err != nil {
				return nil, err
			}
		}
		for _, item := range req.OutExperts {
			newAssign := Assignment{
				DefaultField: field.DefaultField{},
				Uuid:         utils.GenUuidV4(),
				UserId:       item,
				ProjectId:    req.ProjectId,
				Role:         5,
				Operator:     req.Operator,
				IsConfirmed:  false,
				Status:       0,
			}
			_, err := database.MgoAssignments.InsertOne(context.Background(), &newAssign)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	}
	_, err := database.MongoClient.DoTransaction(ctx, callback)
	if err != nil {
		logger.Recorder.Warn("[mongo] make assignments group error: " + err.Error())
		return response.AssignmentCreateFail
	}
	return response.SUCCESS
}

func GetUserAssignments(uid string) (*[]Assignment, int) {
	var assignments []Assignment
	err := database.MgoAssignments.Find(context.Background(), bson.M{
		"user_id": uid,
	}).All(&assignments)
	if err != nil {
		logger.Recorder.Error("[mongo assignment] find user's assign error: " + err.Error())
		return nil, response.AssignmentUserGetFail
	}
	return &assignments, response.SUCCESS
}

func GetProjectAssignment(pid string) (map[string][]Assignment, int) {
	var assignments []Assignment
	err := database.MgoAssignments.Find(context.Background(), bson.M{
		"project_id": pid,
	}).All(&assignments)
	if err != nil {
		logger.Recorder.Error("[mongo assignment] find project's assign error: " + err.Error())
		return nil, response.AssignmentProjectGetFail
	}
	// classify roles
	result := make(map[string][]Assignment)
	result["admins"] = []Assignment{}
	result["experts"] = []Assignment{}
	result["assistants"] = []Assignment{}
	result["teachers"] = []Assignment{}
	result["out_experts"] = []Assignment{}
	for _, assignment := range assignments {
		if assignment.Role == 1 {
			result["admins"] = append(result["admins"], assignment)
			continue
		}
		if assignment.Role == 2 {
			result["experts"] = append(result["experts"], assignment)
			continue
		}
		if assignment.Role == 3 {
			result["assistants"] = append(result["assistants"], assignment)
			continue
		}
		if assignment.Role == 4 {
			result["teachers"] = append(result["teachers"], assignment)
			continue
		}
		if assignment.Role == 5 {
			result["out_experts"] = append(result["out_experts"], assignment)
			continue
		}
	}
	return result, response.SUCCESS
}

func ChangeAssignment(req *request.ChangeAssignment) int {
	err := database.MgoAssignments.UpdateOne(context.Background(),
		bson.M{"uuid": req.AssignmentId},
		bson.M{
			"$set": bson.M{
				"role":         req.NewRole,
				"operator":     req.Operator,
				"is_confirmed": false,
			},
		},
	)
	if err != nil {
		logger.Recorder.Error("[mongo change assign] change role error: " + err.Error())
		return response.AssignmentChangeFail
	}
	return response.SUCCESS
}

func ConfirmAssignment(aid string) int {
	err := database.MgoAssignments.UpdateOne(context.Background(),
		bson.M{"uuid": aid},
		bson.M{
			"$set": bson.M{
				"is_confirmed": true,
			}},
	)
	if err != nil {
		logger.Recorder.Error("[mongo confirm assign] confirm role error: " + err.Error())
		return response.AssignConfirmFail
	}
	return response.SUCCESS
}

func RemoveAssignment(aid string) int {
	err := database.MgoAssignments.Remove(context.Background(), bson.M{"uuid": aid})
	if err != nil {
		logger.Recorder.Error("[mongo delete assign] delete role error: " + err.Error())
		return response.AssignmentDeleteFail
	}
	return response.SUCCESS
}
