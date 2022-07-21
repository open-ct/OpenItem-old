package models

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo/field"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"proj-review/database"
	"proj-review/log"
	"proj-review/request"
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

// DoMakeAssignment
func DoMakeAssignment(makeAssignReq *request.MakeAssignment) (*response.AssignmentDefault, bool) {
	// check user and project status
	user, ok := getUserById(makeAssignReq.UserId)
	if !ok || user.Name == "" {
		return &response.AssignmentDefault{
			Description: "User ID error",
		}, false
	}
	proj, err := getProjectById(makeAssignReq.ProjectId)
	if err != nil {
		return &response.AssignmentDefault{
			Description: "check project",
		}, false
	}
	if proj.Status == 3 {
		return &response.AssignmentDefault{
			Description: "project has been terminated",
		}, false
	}
	// check whether user has been assigned -> set unique index
	newAssignment := Assignment{
		Uuid:        utils.GenUuidV4(),
		UserId:      makeAssignReq.UserId,
		ProjectId:   makeAssignReq.ProjectId,
		Role:        makeAssignReq.Role,
		Operator:    makeAssignReq.Operator,
		IsConfirmed: false,
	}
	result, err := database.MgoAssignments.InsertOne(context.Background(), &newAssignment)
	if err != nil {
		log.Logger.Error("[Mongo Assignment] " + err.Error())
		return &response.AssignmentDefault{
			Description: "create failed",
		}, false
	}
	log.Logger.Info(fmt.Sprintf("Created new assignment: %s", result.InsertedID))
	return &response.AssignmentDefault{
		AssignmentId: newAssignment.Uuid,
		Description:  "ok",
	}, true
}

// DoChangeAssignment
func DoChangeAssignment(changeReq *request.ChangeAssignment) (*response.AssignmentDefault, bool) {
	var orgAssign Assignment
	err := database.MgoAssignments.Find(context.Background(), bson.M{
		"uuid": changeReq.AssignmentId,
	}).One(&orgAssign)
	if err != nil || orgAssign.Status < 0 { // < 0: assignment has been set as invalid.
		return &response.AssignmentDefault{
			Description: "assignment status error",
		}, false
	}
	err = database.MgoAssignments.UpdateOne(context.Background(), bson.M{"uuid": orgAssign.Uuid}, bson.M{
		"$set": bson.M{
			"role":         changeReq.NewRole,
			"operator":     changeReq.Operator,
			"is_confirmed": false, // reset, need to be confirmed again.
		},
	})
	if err != nil {
		log.Logger.Warn("[Assignment change]" + err.Error())
		return &response.AssignmentDefault{
			Description: "Change assignment error",
		}, false
	}
	return &response.AssignmentDefault{
		AssignmentId: orgAssign.Uuid,
		Description:  "ok",
	}, true
}

// DoRemoveAssigment
func DoRemoveAssignment(aid string) (*response.AssignmentDefault, bool) {
	err := database.MgoAssignments.Remove(context.Background(), bson.M{
		"uuid": aid,
	})
	if err != nil {
		log.Logger.Error("[Mongo Remove] " + err.Error())
		return &response.AssignmentDefault{
			AssignmentId: aid,
			Description:  "fail",
		}, false
	}
	return &response.AssignmentDefault{
		AssignmentId: aid,
		Description:  "success",
	}, false
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
