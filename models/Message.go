package models

import (
	"chatapp/config"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	From      string             `json:"from,omitempty" bson:"from,omitempty"`
	To        string             `json:"to,omitempty" bson:"to,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt string             `json:"-"`
}

type SendMessageRequest struct {
	To      string `json:"to,omitempty"`
	Content string `json:"content,omitempty"`
}

type messages []*Message

// contains check if given string is in string array
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// SendMessage create message "from" user to "to" user with given content
func SendMessage(from, to, content string) error {
	messageCollection := config.MI.DB.Collection("messages")
	userCollection := config.MI.DB.Collection("users")

	message := new(Message)
	message.From = from
	message.To = to
	message.Content = content
	message.CreatedAt = time.Now().UTC().String()

	toUser := new(User)
	res := userCollection.FindOne(context.TODO(), bson.M{"username": message.To})

	err := res.Decode(&toUser)
	if err != nil {
		return errors.New("User not found")
	}
	if contains(toUser.BlockedUsers, message.From) {
		return errors.New("you are blocked from this user")
	}

	_, err = messageCollection.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}

	return nil
}

// GetUsersMessagesByUser find and get users messages with by spesific user.
func GetUsersMessagesByUser(username string, with string) (messages, error) {
	messageCollection := config.MI.DB.Collection("messages")
	var results messages

	cur, err := messageCollection.Find(context.TODO(), bson.M{"$or": []bson.M{bson.M{"from": username, "to": with}, bson.M{"to": username, "from": with}}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var msg Message
		err := cur.Decode(&msg)
		if err != nil {
			return nil, err
		}

		results = append(results, &msg)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// GetUsersMessages find all message written by or to user
func GetUsersMessages(username string) (messages, error) {
	messageCollection := config.MI.DB.Collection("messages")
	var results messages

	cur, err := messageCollection.Find(context.TODO(), bson.M{"$or": []bson.M{bson.M{"from": username}, bson.M{"to": username}}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var msg Message
		err := cur.Decode(&msg)
		if err != nil {
			return nil, err
		}

		results = append(results, &msg)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
