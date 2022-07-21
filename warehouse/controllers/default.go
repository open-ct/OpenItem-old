package controllers

import (
	"encoding/json"
	"warehouse/logger"
)

func unmarshalBody(body []byte, obj interface{}) error {
	// marshal json body
	if err := json.Unmarshal(body, obj); err != nil {
		logger.Recorder.Warn("[JSON] " + err.Error())
		return err
	}
	return nil
}
