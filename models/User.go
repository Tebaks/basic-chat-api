package models

import (
	"chatapp/pkg/config"
	"chatapp/pkg/utils"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username,omitempty" bson:"username,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	BlockedUsers []string           `json:"blocked_users,omitempty" bson:"blockedusers,omitempty"`
	CreatedAt    string             `json:"-"`
}

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type BlockRequest struct {
	Username string `json:"username,omitempty"`
}

// BlockUser add blocked user to blocker user's BlockedUsers array
func BlockUser(blocker string, blocked string) error {
	userCollection := config.MI.DB.Collection("users")
	res := userCollection.FindOne(context.TODO(), bson.M{"username": blocker})
	var user User
	if err := res.Decode(&user); err != nil {
		return err
	}
	user.BlockedUsers = append(user.BlockedUsers, blocked)
	filter := bson.M{"username": blocker}
	pushToArray := bson.M{"$set": bson.M{"blockedusers": user.BlockedUsers}}
	result, err := userCollection.UpdateOne(context.TODO(), filter, pushToArray)

	if result.ModifiedCount <= 0 {
		return errors.New("block failed")
	}

	return err
}

// SignUp create new user with hashed password and createdAt timestamp and return token
func SignUp(newUser *User) (string, error) {
	userCollection := config.MI.DB.Collection("users")
	existUser := new(User)
	existUsername := userCollection.FindOne(context.TODO(), bson.M{"username": newUser.Username})
	err := existUsername.Decode(existUser)
	if err == nil {
		return "", errors.New("this username is taken")
	}
	newUser.CreatedAt = time.Now().UTC().String()

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	newUser.Password = string(hash)
	_, err = userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return "", err
	}

	token, err := utils.JWTGenerator(newUser.Username)
	if err != nil {
		return "", err
	}

	return token, nil

}

// Login check for user if user exist create token and return
func Login(credentials *LoginRequest) (string, error) {
	userCollection := config.MI.DB.Collection("users")

	user := new(User)
	existUser := userCollection.FindOne(context.TODO(), bson.M{"username": credentials.Username})
	err := existUser.Decode(user)
	if err != nil {
		return "", errors.New("wrong credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", err
	}

	token, err := utils.JWTGenerator(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
