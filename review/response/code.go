package response

var (
	SUCCESS = 1000 // operation exec successfully

	FAIL = 2000 // operation exec failed with default error

	// Database operations (mongodb)
	DatabaseError = 2001

	// user operations
	UserRegisterFail          = 2100
	UserRegisterPasswordError = 2101
	UserLoginFail             = 2102
	UserAuthError             = 2103
	UserNotExist              = 2104
	UserGenerateTokenError    = 2105
	UserUpdatePasswordError   = 2106
)
