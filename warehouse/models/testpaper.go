package models

import (
	"github.com/qiniu/qmgo/field"
	"time"
)

type QuestionItem struct {
	QuestionId string `json:"question_id" bson:"question_id"`
	Score      int    `json:"score" bson:"score"`
	Comment    string `json:"comment" bson:"comment"`
}

type TestpaperPart struct {
	Title        string         `json:"title" bson:"title"`
	Description  string         `json:"description" bson:"description"`
	QuestionList []QuestionItem `json:"question_list" bson:"question_list"`
	Score        int            `json:"score" bson:"score"`
}

type TestpaperProps struct {
	GradeRange []string `json:"grade_range" bson:"grade_range"`
	Subjects   []string `json:"subjects" bson:"subjects"`
	Difficulty string   `json:"difficulty" bson:"difficulty"`
	TimeLimit  string   `json:"time_limit" bson:"time_limit"`
}

type TestpaperComment struct {
	TimePoint time.Time `json:"time_point" bson:"time_point"`
	Comment   string    `json:"comment" bson:"comment"`
	Author    string    `json:"author" bson:"author"`
}

type TempTestpaper struct {
	field.DefaultField `bson:",inline"`
	Uuid               string             `json:"uuid" bson:"uuid"`
	IsRoot             bool               `json:"is_root" bson:"is_root"`
	Base               string             `json:"base" bson:"base"`
	SourceProject      string             `json:"source_project" bson:"source_project"`
	Author             string             `json:"author" bson:"author"`
	Title              string             `json:"title" bson:"title"`
	Info               []TestpaperPart    `json:"info" bson:"info"`
	Props              TestpaperProps     `json:"props" bson:"props"`
	CommentRecord      []TestpaperComment `json:"comment_record" bson:"comment_record"`
}

type FinalTestpaper struct {
	field.DefaultField `bson:",inline"`
	Uuid               string          `json:"uuid" bson:"uuid"`
	SourceProject      string          `json:"source_project" bson:"source_project"`
	Author             string          `json:"author" bson:"author"`
	Title              string          `json:"title" bson:"title"`
	Info               []TestpaperPart `json:"info" bson:"info"`
	Props              TestpaperProps  `json:"props" bson:"props"`
}
