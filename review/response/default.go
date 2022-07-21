package response

type Default struct {
	OperationCode int         `json:"operation_code"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data"`
}

func GenResponse(opCode int, message string, data interface{}) Default {
	return Default{
		OperationCode: opCode,
		Message:       message,
		Data:          data,
	}
}
