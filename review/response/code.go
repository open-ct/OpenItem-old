package response

var (
	SUCCESS = 1000 // operation exec successfully

	FAIL = 2000 // operation exec failed with default error

	// Database operations (mongodb)
	DatabaseDefaultError       = 2001
	DatabaseInsertError        = 2002
	DatabaseUpdateError        = 2003
	DatabaseDeleteError        = 2004
	DatabaseNoRecord           = 2005
	DatabaseUniqueKeyDuplicate = 2006

	// user operations
	UserRegisterFail          = 2100
	UserRegisterPasswordError = 2101
	UserLoginFail             = 2102
	UserAuthError             = 2103
	UserNotExist              = 2104
	UserGenerateTokenError    = 2105
	UserUpdatePasswordError   = 2106

	// file operations
	FileUploadFail       = 2200
	FileGetFileFail      = 2201
	FileDownloadFileFail = 2202
)
