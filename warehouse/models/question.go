package models

import (
	"github.com/qiniu/qmgo/field"
	"time"
)

// define some question data structure.

type QuestionInfo struct {
	Title    string `json:"title" bson:"title"`
	Type     string `json:"type" bson:"type"`
	Body     string `json:"body" bson:"body"`
	Answer   string `json:"answer" bson:"answer"`
	Solution string `json:"solution" bson:"solution"`
}

type QuestionBasicProps struct {
	Encode              string   `json:"encode" bson:"encode"`
	Subject             string   `json:"subject" bson:"subject"`
	DetailsDimension    string   `json:"details_dimension" bson:"details_dimension"`
	SubDetailsDimension string   `json:"sub_details_dimension" bson:"sub_details_dimension"`
	AbilityDimension    string   `json:"ability_dimension" bson:"ability_dimension"`
	SubAbilityDimension string   `json:"sub_ability_dimension" bson:"sub_ability_dimension"`
	Description         string   `json:"description" bson:"description"`
	SubjectRequirements string   `json:"subject_requirements" bson:"subject_requirements"`
	Details             string   `json:"details" bson:"details"`
	Keywords            []string `json:"keywords" bson:"keywords"`
}

type QuestionSpecProps struct {
	Topic       string `json:"topic" bson:"topic"`
	ArticleType string `json:"article_type" bson:"article_type"`
	Length      string `json:"length" bson:"length"`
}

type QuestionExtraProps struct {
	IsScene              bool   `json:"is_scene" bson:"is_scene"`
	IsQuestionGroup      bool   `json:"is_question_group" bson:"is_question_group"`
	ReadingMaterialTopic string `json:"reading_material_topic" bson:"reading_material_topic"`
	MaterialLength       int    `json:"material_length" bson:"material_length"`
}

type QuestionAdvancedProps struct {
	CttLevel  float64 `json:"ctt_level" bson:"ctt_level"`
	CttDiff_1 float64 `json:"ctt_diff_1" bson:"ctt_diff_1"`
	CttDiff_2 float64 `json:"ctt_diff_2" bson:"ctt_diff_2"`
	IrtLevel  float64 `json:"irt_level" bson:"irt_level"`
}

type QuestionComment struct {
	TimePoint time.Time `json:"time_point" bson:"time_point"`
	Comment   string    `json:"comment" bson:"comment"`
	Author    string    `json:"author" bson:"author"`
}

type QuestionApplyRecord struct {
	GradeFits        string   `json:"grade_fits" bson:"grade_fits"`
	TestYear         string   `json:"test_year" bson:"test_year"`
	TestRegion       []string `json:"test_region" bson:"test_region"`
	ParticipantCount int      `json:"participant_count" bson:"participant_count"`
	TestCount        int      `json:"test_count" bson:"test_count"`
}

type TempQuestion struct {
	field.DefaultField `bson:",inline"`
	Uuid               string                `json:"uuid" bson:"uuid"`
	IsRoot             bool                  `json:"is_root" bson:"is_root"`               // 临时题目是否是根
	Base               string                `json:"base" bson:"base"`                     // 若不是root, 需要设置上级题目, 进行版本管理
	SourceProject      string                `json:"source_project" bson:"source_project"` // 项目来源
	Author             string                `json:"author" bson:"author"`
	Info               QuestionInfo          `json:"info" bson:"info"`
	BasicProps         QuestionBasicProps    `json:"basic_props" bson:"basic_props"`
	SpecProps          QuestionSpecProps     `json:"spec_props" bson:"spec_props"`
	ExtraProps         QuestionExtraProps    `json:"extra_props" bson:"extra_props"`
	AdvancedProps      QuestionAdvancedProps `json:"advanced_props" bson:"advanced_props"`
	ApplyRecord        QuestionApplyRecord   `json:"apply_record" bson:"apply_record"`
	CommentRecord      []QuestionComment     `json:"comment_record" bson:"comment_record"`
}

type FinalQuestion struct {
	field.DefaultField `bson:",inline"`
	Uuid               string                `json:"uuid" bson:"uuid"`
	SourceProject      string                `json:"source_project" bson:"source_project"` // 来源项目id
	FinalVersion       string                `json:"final_version" bson:"final_version"`   // 录入final的最后一个版本
	Author             string                `json:"author" bson:"author"`
	Info               QuestionInfo          `json:"info" bson:"info"`
	BasicProps         QuestionBasicProps    `json:"basic_props" bson:"basic_props"`
	SpecProps          QuestionSpecProps     `json:"spec_props" bson:"spec_props"`
	ExtraProps         QuestionExtraProps    `json:"extra_props" bson:"extra_props"`
	AdvancedProps      QuestionAdvancedProps `json:"advanced_props" bson:"advanced_props"`
	ApplyRecord        QuestionApplyRecord   `json:"apply_record" bson:"apply_record"`
}
