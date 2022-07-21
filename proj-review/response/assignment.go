package response

import "time"

type AssignmentDefault struct {
	AssignmentId string `json:"assignment_id"`
	Description  string `json:"description"`
}

type GetAssignments struct {
	Count       int              `json:"count"`
	Assignments []AssignmentItem `json:"assignments"`
	Description string           `json:"description"`
}

type AssignmentItem struct {
	Uuid        string    `json:"uuid"`
	UserId      string    `json:"user_id"`
	UserName    string    `json:"user_name"`
	ProjectId   string    `json:"project_id"`
	ProjectName string    `json:"project_name"`
	Role        int       `json:"role"`
	IsConfirmed bool      `json:"is_confirmed"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
