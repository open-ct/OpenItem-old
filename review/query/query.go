package query

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"review/database"
	"review/models"
)

// get user list
func QueryUsers(ids []string) map[string]models.User {
	users := make(map[string]models.User)
	for _, id := range ids {
		var u models.User
		err := database.MgoUsers.Find(context.Background(), bson.M{
			"uuid": id,
		}).One(&u)
		if err != nil {
			continue
		}
		users[id] = u
	}
	return users
}

// get project list
func QueryProjects(ids []string) map[string]models.Project {
	projs := make(map[string]models.Project)
	for _, id := range ids {
		var p models.Project
		err := database.MgoProjects.Find(context.Background(), bson.M{
			"uuid": id,
		}).One(&p)
		if err != nil {
			continue
		}
		projs[id] = p
	}
	return projs
}

// get assignment list
func QueryAssignments(ids []string) map[string]models.Assignment {
	assigns := make(map[string]models.Assignment)
	for _, id := range ids {
		var a models.Assignment
		err := database.MgoAssignments.Find(context.Background(), bson.M{
			"uuid": id,
		}).One(&a)
		if err != nil {
			continue
		}
		assigns[id] = a
	}
	return assigns
}

// get file list
func QueryFiles(ids []string) map[string]models.FileItem {
	files := make(map[string]models.FileItem)
	for _, id := range ids {
		var p models.FileItem
		err := database.MgoFileRecords.Find(context.Background(), bson.M{
			"uuid": id,
		}).One(&p)
		if err != nil {
			continue
		}
		files[id] = p
	}
	return files
}

// get step list
func QuerySteps(ids []string) map[string]models.Step {
	steps := make(map[string]models.Step)
	for _, id := range ids {
		var p models.Step
		err := database.MgoSteps.Find(context.Background(), bson.M{
			"uuid": id,
		}).One(&p)
		if err != nil {
			continue
		}
		steps[id] = p
	}
	return steps
}

// get submit list
func QuerySubmits(ids []string) map[string]models.Submit {
	submits := make(map[string]models.Submit)
	for _, id := range ids {
		var s models.Submit
		err := database.MgoSubmits.Find(context.Background(), bson.M{
			"uuid": id,
		}).One(&s)
		if err != nil {
			continue
		}
		submits[id] = s
	}
	return submits
}

// get audit list
func QueryAudits(ids []string) map[string]models.Audit {
	audits := make(map[string]models.Audit)
	for _, id := range ids {
		var a models.Audit
		err := database.MgoAudits.Find(context.Background(), bson.M{
			"uuid": id,
		}).One(&a)
		if err != nil {
			continue
		}
		audits[id] = a
	}
	return audits
}
