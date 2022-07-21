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

type Submit struct {
	field.DefaultField `bson:",inline"`
	Uuid               string    `json:"uuid" bson:"uuid"`
	StepId             string    `json:"step_id" bson:"step_id"`
	Title              string    `json:"title" bson:"title"`
	Description        string    `json:"description" bson:"description"`
	Submitter          string    `json:"submitter" bson:"submitter"`
	Contents           []Content `json:"contents" bson:"contents"`
	Status             int       `json:"status" bson:"status"`
}

type Content struct {
	Uuid    string `json:"uuid" bson:"uuid"`
	Type    int    `json:"type" bson:"type"`
	ItemId  string `json:"item_id" bson:"item_id"`
	Version string `json:"version" bson:"version"`
	Comment string `json:"comment" bson:"comment"`
}

// type: 0-files, 1-question, 2-test paper
// Status: -1-processing, 1-passed, 2-closed

func init() {
	// clean old indexes
	database.MgoSubmits.DropAllIndexes(context.Background())
	err := database.MgoSubmits.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
			{Key: []string{"step_id"}, Unique: false},
			{Key: []string{"Submitter"}, Unique: false},
			//{Key: []string{"contents.uuid"}, Unique: true},
		},
	)
	if err != nil {
		logger.Recorder.Error("[Mongo Submits] " + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in submits collection successfully")
	return
}

func GetOneSubmit(subId string) (*Submit, int) {
	var submit Submit
	err := database.MgoSubmits.Find(context.Background(), bson.M{
		"uuid": subId,
	}).One(&submit)
	if err != nil {
		logger.Recorder.Warn("find submit info err: ", err.Error())
		return nil, response.SubmitGetInfoFail
	}
	return &submit, response.SUCCESS
}

func GetStepSubmits(stepId string) (*[]Submit, int) {
	var submits []Submit
	err := database.MgoSubmits.Find(context.Background(), bson.M{
		"step_id": stepId,
	}).All(&submits)
	if err != nil {
		logger.Recorder.Warn("find submits info err: " + err.Error())
		return nil, response.SubmitGetInfoFail
	}
	return &submits, response.SUCCESS
}

func GetUserSubmitsInStep(req *request.GetUserSubmitsInStep) (*[]Submit, int) {
	var submits []Submit
	err := database.MgoSubmits.Find(context.Background(), bson.M{
		"submitter": req.UserId,
		"step_id":   req.StepId,
	}).All(&submits)
	if err != nil {
		logger.Recorder.Warn("find submits info err: ", err.Error())
		return nil, response.SubmitGetInfoFail
	}
	return &submits, response.SUCCESS
}

func MakeOneSubmit(req *request.CreateSubmit) (*Submit, int) {
	newSubmit := Submit{
		Uuid:        utils.GenUuidV4(),
		StepId:      req.StepId,
		Title:       req.Title,
		Description: req.Description,
		Submitter:   req.UserId,
	}
	insert, err := database.MgoSubmits.InsertOne(context.Background(), &newSubmit)
	if err != nil {
		logger.Recorder.Warn("[mongo create new submit failed] " + err.Error())
		return nil, response.SubmitCreateFail
	}
	logger.Recorder.Info("[Mongo Insert] " + fmt.Sprintf("%s", insert.InsertedID))
	return &newSubmit, response.SUCCESS
}

func AppendContent(req *request.AppendContentInSubmit) (*[]Content, int) {
	var submit Submit
	err := database.MgoSubmits.Find(context.Background(), bson.M{
		"uuid": req.SubmitId,
	}).One(&submit)
	if err != nil {
		logger.Recorder.Warn("Address submit error: " + err.Error())
		return nil, response.SubmitUpdateFail
	}
	contents := submit.Contents
	newContent := Content{
		Uuid:    utils.GenUuidV4(),
		Type:    req.Type,
		ItemId:  req.ItemId,
		Version: req.Version,
		Comment: req.Comment,
	}
	contents = append(contents, newContent)
	err = database.MgoSubmits.UpdateOne(context.Background(),
		bson.M{"uuid": req.SubmitId},
		bson.M{
			"$set": bson.M{
				"contents": contents,
			},
		})
	if err != nil {
		logger.Recorder.Warning("[mongo] append a content error: " + err.Error())
		return nil, response.SubmitUpdateFail
	}
	return &contents, response.SUCCESS
}

func WithdrawContent(req *request.WithdrawContentInSubmit) (*[]Content, int) {
	var submit Submit
	err := database.MgoSubmits.Find(context.Background(), bson.M{
		"uuid": req.SubmitId,
	}).One(&submit)
	if err != nil {
		logger.Recorder.Warn("Address submit error: " + err.Error())
		return nil, response.SubmitUpdateFail
	}
	contents := submit.Contents
	for index, content := range contents {
		if index == req.ContentIndex || content.Uuid == req.ContentId {
			contents = append(contents[:index], contents[index+1:]...)
			break
		}
		if index == len(contents)-1 {
			return &contents, response.SubmitWithdrawFail
		}
	}
	err = database.MgoSubmits.UpdateOne(context.Background(),
		bson.M{"uuid": req.SubmitId},
		bson.M{
			"$set": bson.M{
				"contents": contents,
			},
		})
	if err != nil {
		logger.Recorder.Warning("[mongo] delete content error: " + err.Error())
		return &submit.Contents, response.SubmitWithdrawFail
	}
	return &contents, response.SUCCESS
}

func SetSubmitStatus(req *request.SetSubmitStatus) int {
	err := database.MgoSubmits.UpdateOne(context.Background(),
		bson.M{"uuid": req.SubmitId},
		bson.M{
			"$set": bson.M{
				"status": req.NewStatus,
			},
		})
	if err != nil {
		logger.Recorder.Warning("[mongo] delete content error: " + err.Error())
		return response.SubmitUpdateFail
	}
	return response.SUCCESS
}

func DeleteSubmit(submitId string) int {
	err := database.MgoSubmits.Remove(context.Background(), bson.M{
		"uuid": submitId,
	})
	if err != nil {
		logger.Recorder.Warning("[mongo] delete submit error: " + err.Error())
		return response.SubmitDeleteFail
	}
	// todo: remove audits
	return response.SUCCESS
}
