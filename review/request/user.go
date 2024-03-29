package request

//UserLogin define the user login struct
type UserLogin struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRegister defines create user request format
type UserRegister struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Location string `json:"location"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   bool   `json:"gender"`
	Degree   string `json:"degree"`
	Position string `json:"position"`
	Employer string `json:"employer"`
	Major    string `json:"major"`
}

// UserUpdatePassword define the update password req format 请求更新密码
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
	Location string `json:"location"`
	Degree   string `json:"degree"`
	Position string `json:"position"`
	Employer string `json:"employer"`
	Major    string `json:"major"`
}

// UserSearch define search users struct
type UserSearch struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Degree   string `json:"degree"`
	Position string `json:"position"`
	Employer string `json:"employer"`
	Major    string `json:"major"`
}

// GetUserInfo: just using http-get with url params
// UserLogout: just using http-get with url params
