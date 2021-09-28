package noticer

import "github.com/qiniu/qmgo/field"

// site message

type SiteMessage struct {
	field.DefaultField `bson:",inline"`
	Uuid               string   `json:"uuid" bson:"uuid"`
	Sender             string   `json:"sender" bson:"sender"`
	Receivers          []string `json:"receivers" bson:"receivers"`
	Message            string   `json:"message" bson:"message"`
}
