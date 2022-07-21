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

type Audit struct {
	field.DefaultField `bson:",inline"`
	Uuid               string `json:"uuid" bson:"uuid"`
	SubmitContent      string `json:"submit_content" bson:"submit_content"`
	Result             int    `json:"result" bson:"result"`
	Comment            string `json:"comment" bson:"comment"`
	Auditor            string `json:"auditor" bson:"auditor"`
}

// result: 0-not pass 1-pass, 2-need correct

func init() {
	// clean old indexes
	database.MgoAudits.DropAllIndexes(context.Background())
	err := database.MgoAudits.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
			{Key: []string{"uuid", "submit_content"}, Unique: true},
		},
	)
	if err != nil {
		logger.Recorder.Error("[Mongo Submits] " + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in submits collection successfully")
	return
}

func GetOneAudit(auditId string) (*Audit, int) {
	var audit Audit
	err := database.MgoAudits.Find(context.Background(), bson.M{
		"uuid": auditId,
	}).One(&audit)
	if err != nil {
		logger.Recorder.Warn("find audit info err: " + err.Error())
		return nil, response.AuditGetInfoFail
	}
	return &audit, response.SUCCESS
}

func GetSubmitAudits(submitId string) (*[]Audit, int) {
	var submit Submit
	err := database.MgoSubmits.Find(context.Background(), bson.M{
		"uuid": submitId,
	}).One(&submit)
	if err != nil {
		logger.Recorder.Warn("address submit info err: " + err.Error())
		return nil, response.AuditGetInfoFail
	}
	var audits []Audit
	for _, content := range submit.Contents {
		var audit Audit
		err := database.MgoAudits.Find(context.Background(), bson.M{
			"submit_content": content.Uuid,
		}).One(&audit)
		if err != nil {
			continue
		} else {
			audits = append(audits, audit)
		}
	}
	return &audits, response.SUCCESS
}

func MakeOneAudit(req *request.MakeOneAudit) (*Audit, int) {
	newAudit := Audit{
		DefaultField:  field.DefaultField{},
		Uuid:          utils.GenUuidV4(),
		SubmitContent: req.SubmitContentId,
		Result:        req.Result,
		Comment:       req.Comment,
		Auditor:       req.UserId,
	}
	insert, err := database.MgoAudits.InsertOne(context.Background(), &newAudit)
	if err != nil {
		logger.Recorder.Warn("[mongo create new audit failed] " + err.Error())
		return nil, response.AuditCreateFail
	}
	logger.Recorder.Info("[Mongo Insert] " + fmt.Sprintf("%s", insert.InsertedID))
	return &newAudit, response.SUCCESS
}

func CorrectAudit(req *request.UpdateAudit) (*Audit, int) {
	err := database.MgoAudits.UpdateOne(context.Background(), bson.M{
		"uuid": req.AuditId,
	}, bson.M{
		"$set": bson.M{
			"result":  req.NewResult,
			"comment": req.NewComment,
			"auditor": req.NewAuditor,
		},
	})
	if err != nil {
		logger.Recorder.Warn("[mongo update new audit failed] " + err.Error())
		return nil, response.AuditCorrectFail
	}
	var newAudit Audit
	database.MgoAudits.Find(context.Background(), bson.M{"uuid": req.AuditId}).One(&newAudit)
	return &newAudit, response.SUCCESS
}

func DeleteAudit(auditId string) int {
	// delete in audit collection
	err := database.MgoAudits.Remove(context.Background(), bson.M{
		"uuid": auditId,
	})
	if err != nil {
		return response.AuditDeleteFail
	}
	return response.SUCCESS
}
