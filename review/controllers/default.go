package controllers

import (
	"encoding/json"
	"fmt"
	"review/access"
	"review/logger"
)

// unmarshalBody parse the response body (json type)
func unmarshalBody(body []byte, obj interface{}) error {
	// marshal json body
	if err := json.Unmarshal(body, obj); err != nil {
		logger.Recorder.Warn("[JSON] " + err.Error())
		return err
	}
	return nil
}

// parseUserToken get user id from token string
func parseUserToken(token string) (string, error) {
	//
	fmt.Println("recv token: ", token)
	//
	tokenData, err := access.ParseToken(token)
	if err != nil {
		return "", err
	}
	//
	fmt.Println("recv user: ", tokenData.UserID)
	fmt.Println("recv token time from ", tokenData.IssuedAt, " to ", tokenData.ExpiresAt)
	//
	return tokenData.UserID, nil
}
