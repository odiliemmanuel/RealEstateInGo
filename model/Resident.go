package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resident struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	Email          string             `bson:"email"`
	PhoneNumber    string             `bson:"phone_number"`
	IsEnabled      bool               `bson:"is_enabled"`
	HouseAddress   string             `bson:"house_address"`
	DateRegistered time.Time          `bson:"date_registered"`
}
