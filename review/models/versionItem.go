package models

import "github.com/qiniu/qmgo/field"

// todo: Abstract the versioned object as a base struct.
type VersionItem struct {
	field.DefaultField `bson:",inline"`
	Uuid               string `bson:"uuid"`
	Creator            string `bson:"creator"`
	Updater            string `bson:"updater"`
	Base               string `bson:"base"`
}
