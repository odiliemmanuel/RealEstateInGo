package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GatePass struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Code       string             `bson:"code"`
	StartTime  time.Time          `bson:"startTime"`
	EndTime    time.Time          `bson:"endTime"`
	Visitor    Visitor            `bson:"visitor"`
	PassType   Type               `bson:"passType"`
	IsValid    bool               `bson:"isValid"`
	ResidentId string             `bson:"residentId"`
}
