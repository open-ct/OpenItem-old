package models

import (
	"github.com/qiniu/qmgo/field"
)

type Step struct {
	field.DefaultField `bson:",inline"`
	Uuid               string `json:"uuid" bson:"uuid"`
	Belong             string `json:"belong" bson:"belong"`
	Name               string `json:"name" bson:"name"`
	Status             int    `json:"status" bson:"status"`
	Tasks              string `json:"tasks" bson:"tasks"`
	Description        string `json:"description" bson:"description"`
}
