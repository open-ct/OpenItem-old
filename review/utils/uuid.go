package utils

import "github.com/google/uuid"

/*
	用于生成uuid使用, 这里使用V1(基于timestamp)和V4版本(基于随机数)
*/

// GenUuidV1 generate the uuid version 1 (timestamp)
func GenUuidV1() (string, error) {
	res, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

// GenUuidV4 generate the uuid version 4 (random number)
func GenUuidV4() string {
	res := uuid.New()
	return res.String()
}
