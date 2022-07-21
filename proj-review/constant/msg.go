package constant

var Message = make(map[int]string)

func init() {
	message := make(map[int]string)
	message[SUCCESS] = "ok"
	message[FAIL] = "发生未知错误"

	Message = message
}
