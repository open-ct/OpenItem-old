package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/scrypt"
	"io"
	"proj-review/log"
	"proj-review/request"
	"time"
)

/*
	Using JWT method to authenticate the user (add the user-uuid in request header)
*/

// JWT : HEADER PAYLOAD SIGNATURE
const (
	SecretKEY              string = "JWT-Secret-Key"
	DEFAULT_EXPIRE_SECONDS int    = 60 * 60 * 24 * 7 // default expired 1 week (7 days)
	PasswordHashBytes             = 16
)

// MyCustomClaims This struct is the payload
type MyCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// JwtPayload This struct is the parsing of token payload
type JwtPayload struct {
	UserName  string `json:"user_name"`
	UserID    string `json:"user_id"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

// GenerateToken generate user token
func GenerateToken(loginInfo *request.UserLogin, UserId string, expiredSeconds int) (string, error) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}

	// Create the Claims
	mySigningKey := []byte(SecretKEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	log.Logger.Info("Token will be expired at " + time.Unix(expireAt, 0).String())

	user := *loginInfo
	claims := MyCustomClaims{
		UserId,
		jwt.StandardClaims{
			Issuer:    user.Email,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireAt,
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Signs the token with a secret
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", errors.New("error: failed to generate token")
	}

	return tokenStr, nil
}

// ParseToken Parse the token data (get the operating user id & name)
func ParseToken(tokenString string) (*JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKEY), nil
		})

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		log.Logger.Info(fmt.Sprintf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt))
		log.Logger.Info(fmt.Sprintf("Token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0)))

		return &JwtPayload{
			UserName:  claims.StandardClaims.Issuer,
			UserID:    claims.UserID,
			IssuedAt:  claims.StandardClaims.IssuedAt,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
		}, nil
	} else {
		log.Logger.Info("[JWT]" + err.Error())
		return nil, errors.New("error: failed to validate token")
	}
}

// RefreshToken update token
func RefreshToken(tokenString string) (newTokenString string, err error) {
	// get previous token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKEY), nil
		})

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return "", err
	}

	mySigningKey := []byte(SecretKEY)
	expireAt := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix() //new expired
	newClaims := MyCustomClaims{
		claims.UserID,
		jwt.StandardClaims{
			Issuer:    claims.StandardClaims.Issuer, //name of token issue
			IssuedAt:  time.Now().Unix(),            //time of token issue
			ExpiresAt: expireAt,
		},
	}

	// generate new token with new claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	// sign the token with a secret
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {
		return "", errors.New("error: failed to generate new fresh json web token")
	}

	return tokenStr, nil
}

// GenerateSalt generate salt
func GenerateSalt() (salt string, err error) {
	buf := make([]byte, PasswordHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", errors.New("error: failed to generate user's salt")
	}

	return fmt.Sprintf("%x", buf), nil
}

// GeneratePassHash generate password hash
func GeneratePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, PasswordHashBytes)
	if err != nil {
		return "", errors.New("error: failed to generate password hash")
	}

	return fmt.Sprintf("%x", h), nil
}
