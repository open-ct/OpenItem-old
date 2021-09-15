package models

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"warehouse/database"
	"warehouse/logger"
)

// 进行批量查询的方法
func QueryTempQuestions(idList []string) map[string]TempQuestion {
	questionsMap := make(map[string]TempQuestion)
	for _, id := range idList {
		var question TempQuestion
		err := database.MgoTempQuestions.Find(
			context.Background(),
			bson.M{
				"uuid": id,
			},
		).One(&question)
		if err != nil {
			logger.Recorder.Error("get temp-question error: [" + id + "]" + err.Error())
			continue
		}
		questionsMap[id] = question
	}
	return questionsMap
}

func QueryFinalQuestions(idList []string) map[string]FinalQuestion {
	questionsMap := make(map[string]FinalQuestion)
	for _, id := range idList {
		var question FinalQuestion
		err := database.MgoFinalQuestions.Find(
			context.Background(),
			bson.M{
				"uuid": id,
			},
		).One(&question)
		if err != nil {
			logger.Recorder.Error("get final-question error: [" + id + "]" + err.Error())
			continue
		}
		questionsMap[id] = question
	}
	return questionsMap
}

func QueryTempTestpaper(idList []string) map[string]TempTestpaper {
	testpaperMap := make(map[string]TempTestpaper)
	for _, id := range idList {
		var testPaper TempTestpaper
		err := database.MgoTempTestPaper.Find(
			context.Background(),
			bson.M{
				"uuid": id,
			},
		).One(&testPaper)
		if err != nil {
			logger.Recorder.Error("get final-test-paper error: [" + id + "]" + err.Error())
			continue
		}
		testpaperMap[id] = testPaper
	}
	return testpaperMap
}

func QueryFinalTestpaper(idList []string) map[string]FinalTestpaper {
	testpaperMap := make(map[string]FinalTestpaper)
	for _, id := range idList {
		var testPaper FinalTestpaper
		err := database.MgoFinalTestPaper.Find(
			context.Background(),
			bson.M{
				"uuid": id,
			},
		).One(&testPaper)
		if err != nil {
			logger.Recorder.Error("get final-test-paper error: [" + id + "]" + err.Error())
			continue
		}
		testpaperMap[id] = testPaper
	}
	return testpaperMap
}
