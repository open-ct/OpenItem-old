package models

type QueryListRequest struct {
	IdList []string `json:"id_list"`
}

// Questions
type NewQuestionRequest struct {
	SourceProject string                `json:"source_project"`
	Author        string                `json:"author"`
	Info          QuestionInfo          `json:"info"`
	BasicProps    QuestionBasicProps    `json:"basic_props"`
	SpecProps     QuestionSpecProps     `json:"spec_props"`
	ExtraProps    QuestionExtraProps    `json:"extra_props"`
	AdvancedProps QuestionAdvancedProps `json:"advanced_props"`
	ApplyRecord   QuestionApplyRecord   `json:"apply_record"`
}

type UpdateQuestionRequest struct {
	BaseQuestion     string                `json:"base_question"`
	Author           string                `json:"author"`
	NewInfo          QuestionInfo          `json:"new_info"`
	NewBasicProps    QuestionBasicProps    `json:"new_basic_props"`
	NewSpecProps     QuestionSpecProps     `json:"new_spec_props"`
	NewExtraProps    QuestionExtraProps    `json:"new_extra_props"`
	NewAdvancedProps QuestionAdvancedProps `json:"new_advanced_props"`
	NewApplyRecord   QuestionApplyRecord   `json:"new_apply_record"`
}

type AddQuestionCommentRequest struct {
	QuestionId string `json:"question_id"`
	Comment    string `json:"comment"`
	Author     string `json:"'author'"`
}

// Testpaper
type NewTestpaperRequest struct {
	SourceProject string          `json:"source_project"`
	Author        string          `json:"author"`
	Title         string          `json:"title"`
	Info          []TestpaperPart `json:"info"`
	Props         TestpaperProps  `json:"props"`
}

type UpdateTestpaperRequest struct {
	BaseTestpaper string          `json:"base_testpaper"`
	Author        string          `json:"author"`
	NewTitle      string          `json:"new_title"`
	NewInfo       []TestpaperPart `json:"new_info"`
	NewProps      TestpaperProps  `json:"new_props"`
}

type AddTestpaperCommentRequest struct {
	TestpaperId string `json:"testpaper_id"`
	Comment     string `json:"comment"`
	Author      string `json:"author"`
}
