package response

import "time"

type ProjectDefault struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Description string `json:"description"`
}

type ProjectInfo struct {
	ProjectId          string      `json:"project_id"`
	ProjectName        string      `json:"project_name"`
	ProjectDescription string      `json:"project_description"`
	Requirement        string      `json:"requirement"`
	Target             string      `json:"target"`
	GradeRange         []string    `json:"grade_range"`
	Subjects           []string    `json:"subjects"`
	Summary            string      `json:"summary"`
	Creator            string      `json:"creator"`
	Status             int         `json:"status"`
	Group              interface{} `json:"group"`
	Timetable          interface{} `json:"timetable"`
	Steps              interface{} `json:"steps"`
	Description        string      `json:"description"`
	CreateAt           time.Time   `json:"create_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
}
