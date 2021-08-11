package models

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"proj-review/auth"
	"proj-review/constant"
	"proj-review/database"
	"proj-review/log"
	"proj-review/request"
	"proj-review/response"
	"proj-review/utils"
	"time"
)

// User user data.
type User struct {
	// 基本信息
	Uuid     string `json:"uuid" bson:"uuid"`
	Name     string `json:"name" bson:"name" gorm:"column:name;not null"`
	Email    string `json:"email" bson:"email" gorm:"column:email;unique;not null"`
	Phone    string `json:"phone" bson:"phone" gorm:"column:phone;unique;not null"`
	Gender   bool   `json:"gender" bson:"gender" gorm:"column:gender;not null"`
	Degree   string `json:"degree" bson:"degree" gorm:"column:degree"`
	Position string `json:"position" bson:"position" gorm:"column:position"`
	Employer string `json:"employer" bson:"employer" gorm:"column:employer"`
	Major    string `json:"major" bson:"major" gorm:"column:major"`
	// 系统信息
	Password string `json:"password" bson:"password" gorm:"column:password;not null"`
	Salt     string `json:"salt" bson:"salt" gorm:"column:salt;not null"`
	// 信息更新
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	IsDeleted bool      `json:"is_deleted" bson:"is_deleted"`
}

type UserBasicInfo struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
}

func init() {
	// create the index of user in mongo db.
	err := database.MgoUsers.CreateIndexes(
		context.Background(),
		[]options.IndexModel{
			{Key: []string{"uuid", "email", "phone"}, Unique: true},
			{Key: []string{"name", "major", "position", "employer"}, Unique: false},
		},
	)
	if err != nil {
		log.Logger.Error("[Mongo]" + err.Error())
		return
	}
	log.Logger.Info("[Mongo] Create the index in users collection successfully")
	return
}

// DoUserRegister
func DoUserRegister(registerReq *request.UserRegister) (*response.UserDefault, bool) {
	newUser := User{
		Uuid:      utils.GenUuidV4(),
		Name:      registerReq.Name,
		Email:     registerReq.Email,
		Phone:     registerReq.Phone,
		Gender:    registerReq.Gender,
		Degree:    registerReq.Degree,
		Position:  registerReq.Position,
		Employer:  registerReq.Employer,
		Major:     registerReq.Major,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}
	if isDuplicate(newUser) {
		return &response.UserDefault{
			Description: constant.RegisterMsg.RepeatUser,
		}, false
	}
	// new user, build password
	saltKey, pwdHash, ok := genPasswordHash(registerReq.Password)
	if !ok {
		return &response.UserDefault{
			Description: constant.RegisterMsg.Unknown,
		}, false
	}
	newUser.Password = pwdHash
	newUser.Salt = saltKey
	createResult, err := database.MgoUsers.InsertOne(context.Background(), newUser)
	if err != nil {
		log.Logger.Error("[Mongo] " + err.Error())
		return &response.UserDefault{
			Description: constant.RegisterMsg.Unknown,
		}, false
	}
	log.Logger.Info(fmt.Sprintf("[User] New user created, operation ID: %s", createResult.InsertedID))
	return &response.UserDefault{
		UserID:      newUser.Uuid,
		Name:        newUser.Name,
		Description: constant.RegisterMsg.Ok,
	}, true
}

// DoUserLogin
func DoUserLogin(loginReq *request.UserLogin) (*response.LoginResponse, bool) {
	// 检测用户是否存在
	logUser := User{
		Phone: loginReq.Phone,
		Email: loginReq.Email,
	}
	// 用户是否存在
	user, ok := getUserByContact(logUser)
	if !ok {
		return &response.LoginResponse{
			Description: constant.LoginMsg.UserNotExist,
		}, false
	}
	// 密码认证
	isPass := passwordAuth(loginReq.Password, user)
	if !isPass {
		return &response.LoginResponse{
			UserID:      user.Uuid,
			Description: constant.LoginMsg.AuthError,
		}, false
	}

	// 生成token
	tokenString, err := auth.GenerateToken(loginReq, user.Uuid, 0)
	if err != nil {
		log.Logger.Error("[Login] " + err.Error())
		return &response.LoginResponse{
			UserID:      user.Uuid,
			Description: constant.LoginMsg.Unknown,
		}, false
	}
	// 同时返回用户基本信息
	userInfo := response.UserInfo{
		UserID:   user.Uuid,
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Gender:   user.Gender,
		Degree:   user.Degree,
		Position: user.Position,
		Employer: user.Employer,
		Major:    user.Major,
	}
	return &response.LoginResponse{
		UserID:      user.Uuid,
		UserInfo:    userInfo,
		Token:       tokenString,
		Description: constant.LoginMsg.Ok,
	}, true
}

// DoUpdatePassword
func DoUpdatePassword(updatePwdReq *request.UserUpdatePassword) (*response.UserDefault, bool) {
	var orgUser User
	orgUser, ok := getUserById(updatePwdReq.ID)

	isAuth := passwordAuth(updatePwdReq.OldPassword, orgUser)
	if !isAuth {
		return &response.UserDefault{
			UserID:      orgUser.Uuid,
			Name:        orgUser.Name,
			Description: constant.UpdateInfoMsg.AuthFail,
		}, false
	}
	// 重新生成密码hash
	newSalt, newPwd, ok := genPasswordHash(updatePwdReq.NewPassword)
	if !ok {
		return &response.UserDefault{
			UserID:      orgUser.Uuid,
			Name:        orgUser.Name,
			Description: constant.UpdateInfoMsg.Unknown,
		}, false
	}
	// 更新密码 & salt
	err := database.MgoUsers.UpdateOne(
		context.Background(),
		bson.M{"uuid": updatePwdReq.ID},
		bson.M{
			"$set": bson.M{
				"password": newPwd,
				"salt":     newSalt,
			},
		},
	)
	if err != nil {
		log.Logger.Error("[Mongo] update user password: " + err.Error())
		return &response.UserDefault{
			UserID:      orgUser.Uuid,
			Name:        orgUser.Name,
			Description: constant.UpdateInfoMsg.Unknown,
		}, false
	}
	return &response.UserDefault{
		UserID:      orgUser.Uuid,
		Name:        orgUser.Name,
		Description: constant.UpdateInfoMsg.Ok,
	}, true
}

// DoUpdateUserInfo
func DoUpdateUserInfo(updateInfoReq *request.UserUpdateInfo) (*response.UserDefault, bool) {
	var orgUser User
	orgUser, ok := getUserById(updateInfoReq.ID)
	if !ok {
		return &response.UserDefault{
			UserID:      updateInfoReq.ID,
			Name:        updateInfoReq.Name,
			Description: constant.UpdateInfoMsg.NotExist,
		}, false

	}
	// 检测邮箱 & 电话重复
	isDup := isDuplicate(User{Email: updateInfoReq.Email, Phone: updateInfoReq.Phone})
	if isDup {
		return &response.UserDefault{
			UserID:      updateInfoReq.ID,
			Name:        orgUser.Name,
			Description: constant.UpdateInfoMsg.InfoRepeat,
		}, false
	}
	// todo: 其他信息检查

	err := database.MgoUsers.UpdateOne(context.Background(), bson.M{"uuid": updateInfoReq.ID}, bson.M{
		"$set": bson.M{
			"name":     updateInfoReq.Name,
			"email":    updateInfoReq.Email,
			"phone":    updateInfoReq.Phone,
			"gender":   updateInfoReq.Gender,
			"degree":   updateInfoReq.Degree,
			"position": updateInfoReq.Position,
			"employer": updateInfoReq.Employer,
			"major":    updateInfoReq.Major,
		},
	})
	if err != nil {
		log.Logger.Error("[Mongo] Update user info error: " + err.Error())
		return &response.UserDefault{
			UserID:      updateInfoReq.ID,
			Name:        orgUser.Name,
			Description: constant.UpdateInfoMsg.Unknown,
		}, false
	}
	return &response.UserDefault{
		UserID:      orgUser.Uuid,
		Name:        updateInfoReq.Name,
		Description: constant.UpdateInfoMsg.Ok,
	}, true
}

// DoGetUserInfo
func DoGetUserInfo(userId string) (*response.UserInfo, bool) {
	user, ok := getUserById(userId)
	if !ok {
		return &response.UserInfo{
			UserID:      userId,
			Description: constant.GetInfoMsg.Fail,
		}, false
	}
	userInfo := response.UserInfo{
		Name:        user.Name,
		Email:       user.Email,
		Phone:       user.Phone,
		Gender:      user.Gender,
		Degree:      user.Degree,
		Position:    user.Position,
		Employer:    user.Employer,
		Major:       user.Major,
		Description: constant.GetInfoMsg.Ok,
	}
	return &userInfo, true
}

// DoSearchUsers

// DoDeleteUser

/*
	addition functions
*/

// isUserExist
func isUserExist(userID string) bool {
	count, err := database.MgoUsers.Find(context.Background(), bson.M{
		"uuid": userID,
	}).Count()
	if err != nil || count == 0 {
		log.Logger.Warn("[User] " + err.Error())
		return false
	}
	return true
}

// isDuplicate
func isDuplicate(toCheck User) bool {
	count, err := database.MgoUsers.Find(context.Background(), bson.M{
		"$or": []bson.M{
			bson.M{"email": toCheck.Email},
			bson.M{"phone": toCheck.Phone},
		},
	}).Count()
	if err != nil || count == 0 {
		log.Logger.Warn("[User] " + err.Error())
		return false
	}
	return true
}

// genPasswordHash
func genPasswordHash(password string) (string, string, bool) {
	saltKey, err := auth.GenerateSalt()
	if err != nil {
		log.Logger.Info("[Register] " + err.Error())
		return "", "", false
	}
	pwdHash, err := auth.GeneratePassHash(password, saltKey)
	if err != nil {
		log.Logger.Info("[Register] " + err.Error())
		return "", "", false
	}
	return saltKey, pwdHash, true
}

// passwordAuth
func passwordAuth(pwd string, user User) bool {
	hash, err := auth.GeneratePassHash(pwd, user.Salt)
	if err != nil {
		log.Logger.Error("[Login] " + err.Error())
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
			bson.M{"email": user.Email},
			bson.M{"phone": user.Phone},
		},
	}).One(&u)
	if err != nil || u.Name == "" {
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
	if err != nil || user.Name == "" {
		return User{}, false
	}
	return user, true
}
