package request

type CreateProject struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	Target      string   `json:"target"`
	GradeRange  []string `json:"grade_range"`
	Subjects    []string `json:"subjects"`
	Summary     string   `json:"summary"`
	Creator     string   `json:"creator"`
}

type SearchProject struct {
	Name       string   `json:"name"`
	GradeRange []string `json:"grade_range"`
	Subjects   []string `json:"subjects"`
}

type GetUserProjects struct {
	UserId string `json:"user_id"`
}
