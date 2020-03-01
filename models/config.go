package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ConfigInterface struct {
	ID             primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty" example:"bson id"`
	Title          string              `json:"title"`
	Questions      []QuestionInterface `json:"questions"`
	PublicToken    string              `json:"publicToken"`
	CreatorID      primitive.ObjectID  `json:"creatorId,omitempty"`
	ExpirationDate time.Time           `json:"expirationDate"`
}

type QuestionInterface struct {
	Question   string            `json:"question"`
	AnswerType string            `json:"answerType"`
	Answers    []AnswerInterface `json:"answers"`
	Options    Options           `json:"options,omitempty"`
}

type Options struct {
	Amount     int           `json:"amount"`
	InputType  string        `json:"inputType"`
	Optional   bool          `json:"optional,omitempty"`
	RateScales []interface{} `json:"rateScale,omitempty" swaggertype:"array"`
}

type AnswerInterface struct {
	Text  string      `json:"text"`
	Value interface{} `json:"value" swaggertype:"array"`
}
