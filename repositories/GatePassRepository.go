package repositories

import (
	"RealEstate/config"
	"RealEstate/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GatePassRepository interface {
	SaveGatePass(gatePass models.GatePass) (models.GatePass, error)
	FindAllGatePasses() ([]models.GatePass, error)
	FindGatePassById(gatePassId string) (models.GatePass, error)
	DeleteGatePassById(gatePassId string) error
}

type gatePassRepository struct {
	collection *mongo.Collection
}

func NewGatePassRepository() GatePassRepository {
	return &residentRepository{
		collection: config.DB.Collection("users"),
	}
}

func (r residentRepository) SaveGatePass(gatPass models.GatePass) (models.GatePass, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, gatPass)
	if err != nil {
		return models.GatePass{}, err
	}

	gatPass.ID = result.InsertedID.(primitive.ObjectID)

	return gatPass, nil
}

func (r residentRepository) FindGatePassById(gatePassId string) (models.GatePass, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(gatePassId)
	if err != nil {
		return models.GatePass{}, err
	}

	var gatePass models.GatePass

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&gatePass)
	if err != nil {
		return models.GatePass{}, err
	}

	return gatePass, nil
}

func (r residentRepository) FindAllGatePasses() ([]models.GatePass, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var gatePasses []models.GatePass
	if err := cursor.All(ctx, &gatePasses); err != nil {
		return nil, err
	}
	return gatePasses, nil
}

func (r residentRepository) DeleteGatePassById(gatePasId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()

	// Convert string → ObjectID
	objID, err := primitive.ObjectIDFromHex(gatePasId)
	if err != nil {
		return err
	}

	// Delete
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	// Check if anything was deleted
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
