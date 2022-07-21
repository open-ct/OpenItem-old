package utils

import (
	"encoding/json"
	"review/logger"
)

func StructToJson(v interface{}) string {
	//data, err := json.MarshalIndent(v, "", "  ")
	data, err := json.Marshal(v)
	if err != nil {
		logger.Recorder.Error("convert struct to json error: " + err.Error())
	}
	return string(data)
}

func JsonToStruct(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}
