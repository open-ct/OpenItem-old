package models

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo/field"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"review/access"
	"review/database"
	"review/logger"
	"review/request"
	"review/response"
	"review/utils"
)

// User: define the basic user information of a openct pqbs system user
type User struct {
	field.DefaultField `bson:",inline"`
	// 基本信息
	Uuid    string      `json:"uuid" bson:"uuid"`
	Profile UserProfile `json:"profile" bson:"profile"`
	// 系统信息
	Password string `json:"password" bson:"password"`
	Salt     string `json:"salt" bson:"salt"`
}

type UserProfile struct {
	Name     string `json:"name" bson:"name" `
	Age      int    `json:"age" bson:"age"`
	Locaion  string `json:"locaion" bson:"locaion"`
	Email    string `json:"email" bson:"email" validate:"required"`
	Phone    string `json:"phone" bson:"phone" validate:"required"`
	Gender   bool   `json:"gender" bson:"gender" `
	Degree   string `json:"degree" bson:"degree" `
	Position string `json:"position" bson:"position" `
	Employer string `json:"employer" bson:"employer" `
	Major    string `json:"major" bson:"major"`
}

func init() {
	// create the index of user in mongo db.
	err := database.MgoUsers.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid"}, Unique: true},
			{Key: []string{"profile.email"}, Unique: true},
			{Key: []string{"profile.phone"}, Unique: true},
			{Key: []string{"profile.name"}, Unique: false},
			{Key: []string{"profile.position"}, Unique: false},
			{Key: []string{"profile.employer"}, Unique: false},
			{Key: []string{"profile.major"}, Unique: false},
		},
	)
	if err != nil {
		logger.Recorder.Error("[Mongo]" + err.Error())
		return
	}
	logger.Recorder.Info("[Mongo] Create the index in users collection successfully")
	return
}

func AddUser(req *request.UserRegister) (string, int) {
	newUser := User{
		Uuid: utils.GenUuidV4(),
		Profile: UserProfile{
			Name:     req.Name,
			Age:      0,
			Locaion:  req.Location,
			Email:    req.Email,
			Phone:    req.Phone,
			Gender:   req.Gender,
			Degree:   req.Degree,
			Position: req.Position,
			Employer: req.Employer,
			Major:    req.Major,
		},
	}
	// generate the hash-key & salt
	saltKey, pwdHash, ok := genPasswordHash(req.Password)
	if !ok {
		return "", response.UserRegisterPasswordError
	}
	newUser.Password = pwdHash
	newUser.Salt = saltKey
	createResult, err := database.MgoUsers.InsertOne(context.Background(), &newUser)
	// check duplicate step move to "access"
	if err != nil {
		logger.Recorder.Error("[Mongo] " + err.Error())
		return "", response.DatabaseInsertError
	}
	logger.Recorder.Info(fmt.Sprintf("[User] New user created, operation ID: %s", createResult.InsertedID))
	return newUser.Uuid, response.SUCCESS
}

func UserLogin(req *request.UserLogin) (*User, string, int) {
	fmt.Println("req passwd: ", req.Password)
	toLogin := User{
		Profile: UserProfile{
			Email: req.Email,
			Phone: req.Phone,
		},
	}
	user, ok := getUserByContact(toLogin)
	fmt.Println("user id: ", user.Uuid)
	if !ok {
		return nil, "", response.DatabaseNoRecord
	}
	isPass := passwordAuth(req.Password, user)
	if !isPass {
		return nil, user.Uuid, response.UserAuthError
	}
	// generate token
	token, err := access.GenerateToken(req, user.Uuid, 0)
	if err != nil {
		logger.Recorder.Error("[Login] generate token: " + err.Error())
		return nil, "", response.UserGenerateTokenError
	}
	user.Password = ""
	user.Salt = ""
	return &user, token, response.SUCCESS

}

func GetUser(uid string) (*UserProfile, int) {
	user := User{}
	err := database.MgoUsers.Find(context.Background(), bson.M{
		"uuid": uid,
	}).One(&user)
	if err != nil {
		logger.Recorder.Warning("[Mongo] find user:" + uid + " error")
		return nil, response.DatabaseNoRecord
	}
	return &user.Profile, response.SUCCESS
}

// todo: get users by project id.
func GetProjectUsers() (*[]UserProfile, int) {
	return nil, response.FAIL
}

func UpdateUserInfo(req *request.UserUpdateInfo) (string, int) {
	newInfo := User{
		Profile: UserProfile{
			Name:     req.Name,
			Age:      0,
			Locaion:  req.Location,
			Email:    req.Email,
			Phone:    req.Phone,
			Gender:   req.Gender,
			Degree:   req.Degree,
			Position: req.Position,
			Employer: req.Employer,
			Major:    req.Major,
		},
	}
	// info check ?
	err := database.MgoUsers.UpdateOne(context.Background(), bson.M{"uuid": req.ID}, bson.M{
		"$set": bson.M{
			"profile": newInfo.Profile,
		},
	})
	if err != nil {
		logger.Recorder.Warning("[mongo] update user info: " + err.Error())
		return req.ID, response.DatabaseUpdateError
	}
	return req.ID, response.SUCCESS
}

func UpdateUserPassword(req *request.UserUpdatePassword) (string, int) {
	user, _ := getUserById(req.ID)
	isAuth := passwordAuth(req.OldPassword, user)
	if !isAuth {
		return user.Uuid, response.UserAuthError
	}
	newSalt, newPwd, ok := genPasswordHash(req.NewPassword)
	if !ok {
		return user.Uuid, response.UserUpdatePasswordError
	}
	// 更新密码 & salt
	err := database.MgoUsers.UpdateOne(
		context.Background(),
		bson.M{"uuid": req.ID},
		bson.M{
			"$set": bson.M{
				"password": newPwd,
				"salt":     newSalt,
			},
		},
	)
	if err != nil {
		logger.Recorder.Error("[Mongo] update user password: " + err.Error())
		return user.Uuid, response.DatabaseUpdateError
	}
	return user.Uuid, response.SUCCESS
}

// todo: pages
func SearchUser(req *request.UserSearch) (*[]UserProfile, int) {
	return nil, -1
}

func DeleteUser(uid string) int {
	err := database.MgoUsers.Remove(context.Background(), bson.M{
		"uuid": uid,
	})
	if err != nil {
		logger.Recorder.Error("[User Delete] failed: " + err.Error())
		return response.DatabaseDeleteError
	}
	return response.SUCCESS
}

/*
	Private functions of User
*/

// genPasswordHash: generate users password - hashed
func genPasswordHash(password string) (string, string, bool) {
	saltKey, err := access.GenerateSalt()
	if err != nil {
		logger.Recorder.Info("[Register] " + err.Error())
		return "", "", false
	}
	pwdHash, err := access.GeneratePassHash(password, saltKey)
	if err != nil {
		logger.Recorder.Info("[Register] " + err.Error())
		return "", "", false
	}
	return saltKey, pwdHash, true
}

// passwordAuth
func passwordAuth(pwd string, user User) bool {
	hash, err := access.GeneratePassHash(pwd, user.Salt)
	if err != nil {
		logger.Recorder.Error("[Login] " + err.Error())
		return false
	}
	if hash != user.Password {
		return false
	}
	return true
}

// getUserByContact
func getUserByContact(user User) (User, bool) {
	u := User{}
	err := database.MgoUsers.Find(context.Background(), bson.M{
		"$or": []bson.M{
			bson.M{"profile.email": user.Profile.Email},
			bson.M{"profile.phone": user.Profile.Phone},
		},
	}).One(&u)
	if err != nil || u.Profile.Name == "" {
		return User{}, false
	}
	return u, true
}

// getUserById
func getUserById(userId string) (User, bool) {
	user := User{}
	err := database.MgoUsers.Find(context.Background(), bson.M{
		"uuid": userId,
	}).One(&user)
	if err != nil || user.Profile.Name == "" {
		return User{}, false
	}
	return user, true
}
