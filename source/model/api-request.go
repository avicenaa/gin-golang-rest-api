package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Form struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" swaggerignore:"true"`
	RequestDomain string             `json:"request_domain"`
	RequesterName string             `json:"requester_name"`
	Deadline      string             `json:"deadline"`
	Note          string             `json:"note"`
	Link          string             `json:"link"`
	Tag           []string           `json:"tag"`
	Division      string             `json:"division"`
	InputList     []InputList        `json:"inputList"`
	Assignee      string             `json:"assignee" swaggerignore:"true"`
	Status        string             `json:"status" swaggerignore:"true"`
	ResultAPI     string             `json:"result_api" swaggerignore:"true"`
	CancelNote    string             `json:"cancel_note,omitempty" swaggerignore:"true"`
}

type FormCrawler struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" swaggerignore:"true"`
	Assignee   string             `json:"assignee"`
	Status     string             `json:"status"`
	ResultAPI  string             `json:"result_api"`
	CancelNote string             `json:"cancel_note,omitempty"  swaggerignore:""`
}

type InputList struct {
	Input  []string `json:"input"`
	Output string   `json:"output"`
}
