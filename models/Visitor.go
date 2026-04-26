package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Visitor struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string             `bson:"name"`
	PhoneNumber    string             `bson:"phoneNumber"`
	PurposeOfVisit string             `bson:"purposeOfVisit"`
	ResidentId     string             `bson:"residentId"`
}
