package request

import "time"

/*
	Project
*/

// CreateProject: create new empty project request struct
type CreateProject struct {
	UserId      string   `json:"user_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	Target      string   `json:"target"`
	GradeRange  []string `json:"grade_range"`
	Subjects    []string `json:"subjects"`
	Summary     string   `json:"summary"`
}

// UpdateProjectInfo
type UpdateProjectInfo struct {
	ProjectId   string   `json:"project_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	Target      string   `json:"target"`
	GradeRange  []string `json:"grade_range"`
	Subjects    []string `json:"subjects"`
	Summary     string   `json:"summary"`
}

/*
	Assignment
*/

// MakeOneAssignment
type MakeOneAssignment struct {
	Operator  string `json:"operator"`
	ProjectId string `json:"project_id"`
	UserId    string `json:"user_id"`
	Role      int    `json:"role"`
}

// MakeAssignmentGroup
type MakeAssignmentGroup struct {
	Admins     []string `json:"admin"`
	Experts    []string `json:"expert"`
	Assistants []string `json:"assistant"`
	Teachers   []string `json:"teachers"`
	OutExperts []string `json:"out_experts"`
	ProjectId  string   `json:"project_id"`
	Operator   string   `json:"operator"`
}

// ChangeAssignment
type ChangeAssignment struct {
	Operator     string `json:"operator"`
	AssignmentId string `json:"assignment_id"`
	NewRole      int    `json:"new_role"`
}

// Use Http-GET to RemoveAssignment

/*
	Step
*/

type CreateStep struct {
	Name        string      `json:"name"`
	ProjectId   string      `json:"project_id"`
	Index       int         `json:"index"`
	Description string      `json:"description"`
	Requirement string      `json:"requirement"`
	TimeTable   []TimePoint `json:"time_table"`
	Deadline    int64       `json:"deadline"`
	Creator     string      `json:"creator"`
}

type TimePoint struct {
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Notice    string    `json:"notice"` // todo: noticer config
	Comment   string    `json:"comment"`
}

type AddStepAttachment struct {
	StepId   string   `json:"step_id"`
	FilesIds []string `json:"files_ids"`
	Uploader string   `json:"uploader"`
}

type UpdateStepInfo struct {
	StepId         string `json:"step_id"`
	NewName        string `json:"new_name"`
	NewDeadline    int64  `json:"new_deadline"`
	NewDescription string `json:"new_description"`
	NewRequirement string `json:"new_requirement"`
}

type SetStepStatus struct {
	StepId    string `json:"step_id"`
	NewStatus int    `json:"new_status"`
}

type SetStepTimePoint struct {
	StepId     string `json:"step_id"`
	PointIndex int    `json:"point_index"`
	// index < 0 || index >= len  -> create a new time point
	Info TimePoint `json:"info"`
}

// DeleteTimePoint
type DeleteStepTimePoint struct {
	StepId     string `json:"step_id"`
	PointIndex int    `json:"point_index"`
}

/*
	Submit
*/
type GetUserSubmitsInStep struct {
	UserId string `json:"user_id"`
	StepId string `json:"step_id"`
}

type CreateSubmit struct {
	StepId      string `json:"step_id"`
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AppendContentInSubmit struct {
	SubmitId string `json:"submit_id"`
	Type     int    `json:"type"`
	ItemId   string `json:"item_id"`
	Version  string `json:"version"`
	Comment  string `json:"comment"`
}

type SetSubmitStatus struct {
	SubmitId  string `json:"submit_id"`
	NewStatus int    `json:"new_status"`
}

type WithdrawContentInSubmit struct {
	SubmitId     string `json:"submit_id"`
	ContentIndex int    `json:"content_index"`
	ContentId    string `json:"content_id"`
}

/*
	Audit
*/

type MakeOneAudit struct {
	SubmitContentId string `json:"submit_content_id"`
	Result          int    `json:"result"`
	Comment         string `json:"comment"`
	UserId          string `json:"user_id"`
}

type UpdateAudit struct {
	AuditId    string `json:"audit_id"`
	NewResult  int    `json:"new_result"`
	NewComment string `json:"new_comment"`
	NewAuditor string `json:"new_auditor"`
}
