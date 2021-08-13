package response

// UserDefault
/*
	This response used in:
*/
type UserDefault struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UserInfo
type UserInfo struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Gender      bool   `json:"gender"`
	Degree      string `json:"degree"`
	Position    string `json:"position"`
	Employer    string `json:"employer"`
	Major       string `json:"major"`
	Description string `json:"description"`
}

type userInfo struct {
}

// LoginResponse defines login result response
type LoginResponse struct {
	UserID      string   `json:"user_id"`
	UserInfo    UserInfo `json:"user_info"`
	Token       string   `json:"token"`
	Description string   `json:"description"`
}

// CheckDuplicate
type CheckDuplicate struct {
	Result      bool   `json:"result"`
	Description string `json:"description"`
}
