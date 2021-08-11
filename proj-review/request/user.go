package request

// UserLogin define the user login struct
type UserLogin struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRegister defines create user request format
type UserRegister struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   bool   `json:"gender"`
	Degree   string `json:"degree"`
	Position string `json:"position"`
	Employer string `json:"employer"`
	Major    string `json:"major"`
}

// UserUpdatePassword define the update password req format
type UserUpdatePassword struct {
	ID          string `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// UserUpdateInfo define the user info updating request format
type UserUpdateInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   bool   `json:"gender"`
	Degree   string `json:"degree"`
	Position string `json:"position"`
	Employer string `json:"employer"`
	Major    string `json:"major"`
}

// GetUserInfo: just using http-get with url params
// UserLogout: just using http-get with url params

// todo:
// UserSearch define search users
type UserSearch struct {
}
