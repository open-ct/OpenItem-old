package models

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo/field"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"warehouse/database"
	"warehouse/logger"
	"warehouse/utils"
)

// init mongo database

func initTempQuestion() {
	// drop original indexes
	database.MgoTempQuestions.DropAllIndexes(context.Background())
	err := database.MgoTempQuestions.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
		},
	)
	if err != nil {
		logger.Recorder.Error("[mongo temp question initial] create indexes error: " + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in temp-question collection successfully")
	return
}

func initFinalQuestion() {
	// drop original indexes
	database.MgoFinalQuestions.DropAllIndexes(context.Background())
	err := database.MgoFinalQuestions.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
		},
	)
	if err != nil {
		logger.Recorder.Error("[mongo final question initial] create indexes error: " + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in final-question collection successfully")
	return
}

func initTempTestpaper() {
	// drop original indexes
	database.MgoTempTestPaper.DropAllIndexes(context.Background())
	err := database.MgoTempTestPaper.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
		},
	)
	if err != nil {
		logger.Recorder.Error("[mongo temp test-paper initial] create indexes error: " + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in temp-test-paper collection successfully")
	return
}

func initFinalTestpaper() {
	// drop original indexes
	database.MgoFinalTestPaper.DropAllIndexes(context.Background())
	err := database.MgoFinalTestPaper.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
		},
	)
	if err != nil {
		logger.Recorder.Error("[mongo temp final-paper initial] create indexes error: " + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in final-test-paper collection successfully")
	return
}

func init() {
	initTempQuestion()
	initFinalQuestion()
	initTempTestpaper()
	initFinalTestpaper()
	// initialize over.
}

/**
Questions Operations:
*/
func CreateNewTempQuestion(request *NewQuestionRequest) (string, error) {
	newTempQuestion := TempQuestion{
		Uuid:          utils.GenUuidV4(),
		IsRoot:        true,
		Base:          "root",
		SourceProject: request.SourceProject,
		Author:        request.Author,
		Info:          request.Info,
		BasicProps:    request.BasicProps,
		SpecProps:     request.SpecProps,
		ExtraProps:    request.ExtraProps,
		AdvancedProps: request.AdvancedProps,
		ApplyRecord:   request.ApplyRecord,
		CommentRecord: nil,
	}
	insert, err := database.MgoTempQuestions.InsertOne(context.Background(), &newTempQuestion)
	if err != nil {
		logger.Recorder.Error("insert new temp-question error: " + err.Error())
		return "", err
	}
	logger.Recorder.Info(fmt.Sprintf("new temp-question created: %s", insert.InsertedID))
	return newTempQuestion.Uuid, nil
}

func UpdateQuestion(request *UpdateQuestionRequest) (string, error) {
	var oldQuestion TempQuestion
	err := database.MgoTempQuestions.Find(context.Background(), bson.M{
		"uuid": request.BaseQuestion,
	}).One(&oldQuestion)
	if err != nil {
		return "", err
	}
	newTempQuestion := TempQuestion{
		Uuid:          utils.GenUuidV4(),
		IsRoot:        false,
		Base:          request.BaseQuestion,
		SourceProject: oldQuestion.SourceProject,
		Author:        request.Author,
		Info:          request.NewInfo,
		BasicProps:    request.NewBasicProps,
		SpecProps:     request.NewSpecProps,
		ExtraProps:    request.NewExtraProps,
		AdvancedProps: request.NewAdvancedProps,
		ApplyRecord:   request.NewApplyRecord,
		CommentRecord: nil,
	}
	insert, err := database.MgoTempQuestions.InsertOne(context.Background(), &newTempQuestion)
	if err != nil {
		logger.Recorder.Error("update a temp-question error: " + err.Error())
		return "", err
	}
	logger.Recorder.Info(fmt.Sprintf("temp-question updated: %s", insert.InsertedID))
	return newTempQuestion.Uuid, nil
}

func TraceQuestionVersion(qid string) ([]TempQuestion, error) {
	var endPointQuestion TempQuestion
	err := database.MgoTempQuestions.Find(context.Background(), bson.M{
		"uuid": qid,
	}).One(&endPointQuestion)
	if err != nil {
		logger.Recorder.Error("find base questions failed, qid: [" + qid + "] " + err.Error())
		return nil, err
	}
	var questions []TempQuestion
	questions = append(questions, endPointQuestion)
	isEnd := endPointQuestion.IsRoot
	currentBaseId := endPointQuestion.Base
	for !isEnd {
		var currentNode TempQuestion
		err := database.MgoTempQuestions.Find(context.Background(), bson.M{
			"uuid": currentBaseId,
		}).One(&currentNode)
		if err != nil {
			logger.Recorder.Error("find middle-node questions failed, qid: [" + currentBaseId + "] " + err.Error())
			return questions, err
		}
		questions = append(questions, currentNode)
		isEnd = currentNode.IsRoot == true
		currentBaseId = currentNode.Base
	}
	return questions, nil
}

func AddQuestionComment(request *AddQuestionCommentRequest) error {
	newComment := QuestionComment{
		TimePoint: time.Now(),
		Comment:   request.Comment,
		Author:    request.Author,
	}
	var commentQuestion TempQuestion
	err := database.MgoTempQuestions.Find(context.Background(), bson.M{
		"uuid": request.QuestionId,
	}).One(&commentQuestion)
	if err != nil {
		logger.Recorder.Error("cannot address the question: " + request.QuestionId + " for " + err.Error())
		return err
	}
	newComments := append(commentQuestion.CommentRecord, newComment)
	err = database.MgoTempQuestions.UpdateOne(context.Background(),
		bson.M{
			"uuid": request.QuestionId,
		},
		bson.M{
			"comment_record": newComments,
		},
	)
	if err != nil {
		logger.Recorder.Error("add new comment error: " + err.Error())
		return err
	}
	return nil
}

func FinishTempQuestion(qid string) (string, error) {
	var tempQuestion TempQuestion
	err := database.MgoTempQuestions.Find(context.Background(), bson.M{
		"uuid": qid,
	}).One(&tempQuestion)
	if err != nil {
		logger.Recorder.Error("cannot address the question: " + qid + " for " + err.Error())
		return "", err
	}
	finalQuestion := FinalQuestion{
		Uuid:          utils.GenUuidV4(),
		SourceProject: tempQuestion.SourceProject,
		FinalVersion:  tempQuestion.Uuid,
		Author:        tempQuestion.Author,
		Info:          tempQuestion.Info,
		BasicProps:    tempQuestion.BasicProps,
		SpecProps:     tempQuestion.SpecProps,
		ExtraProps:    tempQuestion.ExtraProps,
		AdvancedProps: tempQuestion.AdvancedProps,
		ApplyRecord:   tempQuestion.ApplyRecord,
	}
	insert, err := database.MgoFinalQuestions.InsertOne(context.Background(), &finalQuestion)
	if err != nil {
		logger.Recorder.Error("conver to final-question failed: " + err.Error())
		return "", err
	}
	logger.Recorder.Info(fmt.Sprintf("convert to final successfully: %s", insert.InsertedID))
	return finalQuestion.Uuid, nil
}

func GetUserTempQuestions(uid string) ([]TempQuestion, error) {
	questions := []TempQuestion{}
	err := database.MgoTempQuestions.Find(
		context.Background(),
		bson.M{
			"author": uid,
		},
	).All(&questions)
	if err != nil {
		logger.Recorder.Error("find user's temp-question error: " + err.Error())
		return nil, err
	}
	return questions, nil
}

func GetUserFinalQuestions(uid string) ([]FinalQuestion, error) {
	questions := []FinalQuestion{}
	err := database.MgoFinalQuestions.Find(
		context.Background(),
		bson.M{
			"author": uid,
		},
	).All(&questions)
	if err != nil {
		logger.Recorder.Error("find user's final-question error: " + err.Error())
		return nil, err
	}
	return questions, nil
}

func GetProjectTempQuestions(pid string) ([]TempQuestion, error) {
	questions := []TempQuestion{}
	err := database.MgoTempQuestions.Find(
		context.Background(),
		bson.M{
			"source_project": pid,
		},
	).All(&questions)
	if err != nil {
		logger.Recorder.Error("find project's temp-question error: " + err.Error())
		return nil, err
	}
	return questions, nil
}

func GetProjectFinalQuestions(pid string) ([]FinalQuestion, error) {
	questions := []FinalQuestion{}
	err := database.MgoFinalQuestions.Find(
		context.Background(),
		bson.M{
			"source_project": pid,
		},
	).All(&questions)
	if err != nil {
		logger.Recorder.Error("find project's final-question error: " + err.Error())
		return nil, err
	}
	return questions, nil
}

/**
TestPaper Operations:
*/
func CreateNewTestpaper(request *NewTestpaperRequest) (string, error) {
	newTestPaper := TempTestpaper{
		Uuid:          utils.GenUuidV4(),
		IsRoot:        true,
		Base:          "root",
		SourceProject: request.SourceProject,
		Author:        request.Author,
		Title:         request.Title,
		Info:          request.Info,
		Props:         request.Props,
	}
	insert, err := database.MgoTempTestPaper.InsertOne(context.Background(), &newTestPaper)
	if err != nil {
		logger.Recorder.Error("insert new temp-tes-paper error: " + err.Error())
		return "", err
	}
	logger.Recorder.Info(fmt.Sprintf("new temp-test-paper created: %s", insert.InsertedID))
	return newTestPaper.Uuid, nil
}

func UpdateTestpaper(request *UpdateTestpaperRequest) (string, error) {
	var oldTestPaper TempTestpaper
	err := database.MgoTempTestPaper.Find(context.Background(), bson.M{
		"uuid": request.BaseTestpaper,
	}).One(&oldTestPaper)
	if err != nil {
		logger.Recorder.Error("base test-paper cannot find: " + err.Error())
		return "", err
	}
	updatedTestPaper := TempTestpaper{
		DefaultField:  field.DefaultField{},
		Uuid:          utils.GenUuidV4(),
		IsRoot:        false,
		Base:          request.BaseTestpaper,
		SourceProject: oldTestPaper.SourceProject,
		Author:        request.Author,
		Title:         request.NewTitle,
		Info:          request.NewInfo,
		Props:         request.NewProps,
	}
	insert, err := database.MgoTempTestPaper.InsertOne(context.Background(), &updatedTestPaper)
	if err != nil {
		logger.Recorder.Error("updated temp-tes-paper error: " + err.Error())
		return "", err
	}
	logger.Recorder.Info(fmt.Sprintf("temp-test-paper updated: %s", insert.InsertedID))
	return updatedTestPaper.Uuid, nil
}

func AddTestpaperComment(requset *AddTestpaperCommentRequest) error {
	newComment := TestpaperComment{
		TimePoint: time.Now(),
		Comment:   requset.Comment,
		Author:    requset.Author,
	}
	var commentTestPaper TempTestpaper
	err := database.MgoTempTestPaper.Find(context.Background(), bson.M{
		"uuid": requset.TestpaperId,
	}).One(&commentTestPaper)
	if err != nil {
		logger.Recorder.Error("cannot address the test-paper: " + err.Error())
		return nil
	}
	newComments := append(commentTestPaper.CommentRecord, newComment)
	err = database.MgoTempTestPaper.UpdateOne(context.Background(),
		bson.M{
			"uuid": requset.TestpaperId,
		},
		bson.M{
			"comment_record": newComments,
		},
	)
	if err != nil {
		logger.Recorder.Error("add new comment error: " + err.Error())
		return err
	}
	return nil
}

func TraceTestpaperVersion(tid string) ([]TempTestpaper, error) {
	var endPointTestPaper TempTestpaper
	err := database.MgoTempTestPaper.Find(context.Background(), bson.M{
		"uuid": tid,
	}).One(&endPointTestPaper)
	if err != nil {
		logger.Recorder.Error("cannot get the end test-paper: " + err.Error())
		return nil, err
	}
	var testPapers []TempTestpaper
	testPapers = append(testPapers, endPointTestPaper)
	isEnd := endPointTestPaper.IsRoot
	currentBaseId := endPointTestPaper.Base
	for !isEnd {
		var currentNode TempTestpaper
		err := database.MgoTempTestPaper.Find(context.Background(), bson.M{
			"uuid": currentBaseId,
		}).One(&currentNode)
		if err != nil {
			logger.Recorder.Error("find middle node cannot find: " + err.Error())
			return testPapers, err
		}
		testPapers = append(testPapers, currentNode)
		currentBaseId = currentNode.Base
		isEnd = currentNode.IsRoot == true
	}
	return testPapers, nil
}

func FinishTempTestpaper(tid string) (string, error) {
	var finishedTestPaper TempTestpaper
	err := database.MgoTempTestPaper.Find(context.Background(), bson.M{
		"uuid": tid,
	}).One(&finishedTestPaper)
	if err != nil {
		logger.Recorder.Error("error to find finished error: " + err.Error())
		return "", nil
	}
	newFinalPaper := FinalTestpaper{
		Uuid:          utils.GenUuidV4(),
		SourceProject: finishedTestPaper.SourceProject,
		Author:        finishedTestPaper.Author,
		Title:         finishedTestPaper.Title,
		Info:          finishedTestPaper.Info,
		Props:         finishedTestPaper.Props,
	}
	insert, err := database.MgoFinalTestPaper.InsertOne(context.Background(), &newFinalPaper)
	if err != nil {
		logger.Recorder.Error("conver to final-test-paper failed: " + err.Error())
		return "", err
	}
	logger.Recorder.Info(fmt.Sprintf("convert to final successfully: %s", insert.InsertedID))
	return newFinalPaper.Uuid, nil
}

func GetUserTempTestpaper(uid string) ([]TempTestpaper, error) {
	testPapers := []TempTestpaper{}
	err := database.MgoTempTestPaper.Find(
		context.Background(),
		bson.M{
			"author": uid,
		},
	).All(&testPapers)
	if err != nil {
		logger.Recorder.Error("find user's temp-test-paper error: " + err.Error())
		return nil, err
	}
	return testPapers, nil
}

func GetUserFinalTestpaper(uid string) ([]FinalTestpaper, error) {
	testPapers := []FinalTestpaper{}
	err := database.MgoFinalTestPaper.Find(
		context.Background(),
		bson.M{
			"author": uid,
		},
	).All(&testPapers)
	if err != nil {
		logger.Recorder.Error("find user's final-test-paper error: " + err.Error())
		return nil, err
	}
	return testPapers, nil
}

func GetProjectTempTestpaper(pid string) ([]TempTestpaper, error) {
	testPapers := []TempTestpaper{}
	err := database.MgoTempTestPaper.Find(
		context.Background(),
		bson.M{
			"source_project": pid,
		},
	).All(&testPapers)
	if err != nil {
		logger.Recorder.Error("find project's temp-test-paper error: " + err.Error())
		return nil, err
	}
	return testPapers, nil
}

func GetProjecgtFinalTestpaper(pid string) ([]FinalTestpaper, error) {
	testPapers := []FinalTestpaper{}
	err := database.MgoFinalTestPaper.Find(
		context.Background(),
		bson.M{
			"source_project": pid,
		},
	).All(&testPapers)
	if err != nil {
		logger.Recorder.Error("find project's final-test-paper error: " + err.Error())
		return nil, err
	}
	return testPapers, nil
}
