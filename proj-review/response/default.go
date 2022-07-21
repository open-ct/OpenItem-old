package response

// Default define the default response format
type Default struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// GenResponse function to generate default response struct.
func GenResponse(code int, msg string, data interface{}) Default {
	return Default{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
