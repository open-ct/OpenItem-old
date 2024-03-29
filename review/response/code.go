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

	// project operations
	ProjectCreateFail  = 2301
	ProjectNotExist    = 2302
	ProjectGetInfoFail = 2303
	ProjectUpdateFail  = 2304

	// project-assignment operations
	AssignmentCreateFail     = 2310
	AssignmentUserGetFail    = 2311
	AssignmentProjectGetFail = 2312
	AssignmentChangeFail     = 2313
	AssignmentDeleteFail     = 2314
	AssignConfirmFail        = 2315

	// project-step operations
	StepCreateFail          = 2320
	StepAddAttachmentsFail  = 2321
	StepUpdateInfoFail      = 2322
	StepChangeStatusFail    = 2323
	StepGetInfoFail         = 2324
	StepSetTimePointFail    = 2325
	StepDeleteTimePointFail = 2326
	StepDeleteFail          = 2327
	StepGetStatisticFail    = 2328

	// project-submit operations
	SubmitGetInfoFail  = 2330
	SubmitCreateFail   = 2331
	SubmitUpdateFail   = 2332
	SubmitWithdrawFail = 2333
	SubmitDeleteFail   = 2334

	// project-audit operations
	AuditGetInfoFail = 2340
	AuditCreateFail  = 2341
	AuditCorrectFail = 2342
	AuditDeleteFail  = 2343
)
