package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type StateInterface struct {
	Answers             map[int]Answer     `json:"answers"`
	ActiveQuestionIndex int                `json:"activeQuestionIndex"`
	Completed           bool               `json:"completed"`
	ConfigID            primitive.ObjectID `json:"-"`
	UserID              primitive.ObjectID `json:"-"`
	ID                  primitive.ObjectID `json:"-"`
}

type Answer struct {
	AnswerInterface `json:"value"`
	Index           int `json:"index"`
}
