package request

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
