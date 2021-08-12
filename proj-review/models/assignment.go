package models

import (
	"context"
	"github.com/qiniu/qmgo/field"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"proj-review/database"
	"proj-review/log"
	"proj-review/response"
	"proj-review/utils"
)

type Assignment struct {
	field.DefaultField `bson:",inline"`
	Uuid               string `bson:"uuid"`
	UserId             string `bson:"user_id"`
	ProjectId          string `bson:"project_id"`
	Role               int    `bson:"role"`
	Operator           string `bson:"operator"`
	IsConfirmed        bool   `bson:"is_confirmed"`
	Status             int    `bson:"status"`
}

// role: 0-admin, 1-expert, 2-assistant, 3-teachers, 4-outer
// Operator "system" or uuid of user

func init() {
	// create index
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
		log.Logger.Error("[Mongo Assignment] " + err.Error())
		return
	}
	log.Logger.Info("[Mongo] Create the index in assignments collection successfully")
	return
}

// createOneAssignment
func createOneAssignment(user string, project string, role int, opr string, confirm bool) error {
	newAssign := Assignment{
		Uuid:        utils.GenUuidV4(),
		UserId:      user,
		ProjectId:   project,
		Role:        role,
		Operator:    opr,
		Status:      0,
		IsConfirmed: confirm,
	}
	_, err := database.MgoAssignments.InsertOne(context.Background(), &newAssign)
	if err != nil {
		return err
	}
	return nil
}

// convertAssignmentToResponse
func convertAssignmentToResponse(a Assignment) response.AssignmentItem {
	var u User
	var p Project
	database.MgoUsers.Find(context.Background(), bson.M{
		"uuid": a.UserId,
	}).One(&u)
	database.MgoProjects.Find(context.Background(), bson.M{
		"uuid": a.ProjectId,
	}).One(&p)

	return response.AssignmentItem{
		Uuid:        a.Uuid,
		UserId:      a.UserId,
		UserName:    u.Name,
		ProjectId:   a.ProjectId,
		ProjectName: p.BasicInfo.Name,
		Role:        a.Role,
		IsConfirmed: a.IsConfirmed,
		Status:      a.Status,
		CreatedAt:   a.CreateAt,
	}
}

// DoGetUserAssignments
func DoGetUserAssignments(userId string) (*response.GetAssignments, bool) {
	var assignments []Assignment
	database.MgoAssignments.Find(context.Background(), bson.M{
		"user_id": userId,
	}).Sort("createAt").All(&assignments)
	if len(assignments) == 0 {
		return &response.GetAssignments{
			Count:       0,
			Description: "no record...",
		}, false
	}
	var assignItems []response.AssignmentItem
	for _, a := range assignments {
		assignItems = append(assignItems, convertAssignmentToResponse(a))
	}
	return &response.GetAssignments{
		Count:       len(assignments),
		Description: "ok",
		Assignments: assignItems,
	}, true
}

// DoGetProjectAssignments
func DoGetProjectAssignments(projId string) (*response.GetAssignments, bool) {
	var assignments []Assignment
	database.MgoAssignments.Find(context.Background(), bson.M{
		"project_id": projId,
	}).Sort("createAt").All(&assignments)
	if len(assignments) == 0 {
		return &response.GetAssignments{
			Count:       0,
			Description: "no record...",
		}, false
	}
	var assignItems []response.AssignmentItem
	for _, a := range assignments {
		assignItems = append(assignItems, convertAssignmentToResponse(a))
	}
	return &response.GetAssignments{
		Count:       len(assignments),
		Description: "ok",
		Assignments: assignItems,
	}, true
}
