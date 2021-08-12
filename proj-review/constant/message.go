package constant

// 向前端传递基本消息响应

type BasicMessage struct {
	// 用户操作 提示
	LoginSuccess      string
	LoginFail         string
	LoginStatusOk     string
	LoginStatusError  string
	RegisterSuccess   string
	RegisterFail      string
	GetInfoSuccess    string
	GetInfoFail       string
	UpdateInfoSuccess string
	UpdateInfoFail    string
	DeleteUserSuccess string
	DeleteUserFail    string

	// 项目操作提示
	CreateProjectSuccess    string
	CreateProjectFail       string
	UpdateProjectSuccess    string
	UpdateProjectFail       string
	TerminateProjectSuccess string
	TerminateProjectFail    string
	FinishProjectSuccess    string
	FinishProjectFail       string

	// 文件操作提示
	UploadFileSuccess   string
	UploadFileFail      string
	DownloadFileSuccess string
	DownloadFileFail    string
	GetFileInfoSuccess  string
	GetFileInfoFail     string
	DeleteFileSuccess   string
	DeleteFileFail      string
}

var BasicMsg = BasicMessage{
	// 用户操作 消息反馈
	LoginSuccess:      "登录成功",
	LoginFail:         "登陆失败",
	LoginStatusOk:     "登录状态正常",
	LoginStatusError:  "登录状态失效，请重新登录",
	RegisterSuccess:   "注册成功",
	RegisterFail:      "注册失败",
	GetInfoSuccess:    "获取信息成功",
	GetInfoFail:       "获取信息失败",
	UpdateInfoSuccess: "信息更新成功",
	UpdateInfoFail:    "信息更新失败",
	DeleteUserFail:    "用户删除失败",
	DeleteUserSuccess: "用户删除成功",

	// 项目操作消息反馈
	CreateProjectSuccess:    "项目创建成功",
	CreateProjectFail:       "项目创建失败",
	UpdateProjectSuccess:    "项目信息更新成功",
	UpdateProjectFail:       "项目信息更新失败",
	TerminateProjectFail:    "中止项目失败",
	TerminateProjectSuccess: "项目已终止",
	FinishProjectSuccess:    "项目已标记完成",
	FinishProjectFail:       "完成项目失败",

	// 文件操作项目反馈
	UploadFileFail:      "上传文件失败",
	UploadFileSuccess:   "上传文件成功",
	DownloadFileFail:    "下载文件失败",
	DownloadFileSuccess: "下载文件成功",
	GetFileInfoFail:     "获取文件信息失败",
	GetFileInfoSuccess:  "获取文件信息成功",
	DeleteFileSuccess:   "文件删除成功",
	DeleteFileFail:      "文件删除失败",
}

// 用户框架 各个操作的详细消息反馈

type RegisterMessage struct {
	Ok          string
	RepeatUser  string
	InvalidInfo string
	Unknown     string
}

var RegisterMsg = RegisterMessage{
	Ok:          "ok",
	RepeatUser:  "用户已存在（手机号或邮箱已被注册）",
	InvalidInfo: "信息不合法",
	Unknown:     "其他错误",
}

type LoginMessage struct {
	Ok           string
	UserNotExist string
	AuthError    string
	Unknown      string
}

var LoginMsg = LoginMessage{
	Ok:           "ok",
	UserNotExist: "用户不存在,请注册",
	AuthError:    "认证错误,请重试",
	Unknown:      "其他错误",
}

type UpdateInfoMessage struct {
	Ok          string
	Unknown     string
	NotExist    string
	InfoRepeat  string
	InvalidInfo string
	AuthFail    string
}

var UpdateInfoMsg = UpdateInfoMessage{
	Ok:          "ok",
	InfoRepeat:  "信息已被注册，请重试",
	NotExist:    "对应信息不存在",
	InvalidInfo: "信息不合法",
	AuthFail:    "认证失败,请重试",
	Unknown:     "其他错误",
}

type GetInfoMessage struct {
	Ok      string
	Fail    string
	Unknown string
}

var GetInfoMsg = GetInfoMessage{
	Ok:      "ok",
	Fail:    "获取信息失败",
	Unknown: "其他错误",
}

type DeleteUserMessage struct {
	Ok      string
	Fail    string
	Unknown string
}

var DeleteUserMsg = DeleteUserMessage{
	Ok:      "ok",
	Fail:    "获取信息失败",
	Unknown: "其他错误",
}

// 项目框架 消息反馈
// todo: 完善项目消息反馈
type ProjectMessage struct {
	Ok      string
	Fail    string
	Unknown string
}

var ProjectMsg = ProjectMessage{
	Ok:      "ok",
	Fail:    "获取信息失败",
	Unknown: "其他错误",
}

// 文件操作 消息反馈
// todo:
type FileMessage struct {
	Ok      string
	Fail    string
	Unknown string
}

var FileMsg = FileMessage{
	Ok:      "ok",
	Fail:    "文件操作失败",
	Unknown: "其他错误",
}
