package request

type MakeAssignment struct {
	UserId    string `json:"user_id"`
	ProjectId string `json:"project_id"`
	Role      int    `json:"role"`
	Operator  string `json:"operator"`
}

type ChangeAssignment struct {
	AssignmentId string `json:"assignment_id"`
	NewRole      int    `json:"new_role"`
	Operator     string `json:"operator"`
}
