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

type ResidentRepository interface {
	Save(entry models.Resident) (models.Resident, error)
	FindById(residentId string) (models.Resident, error)
	FindByEmail(email string) (models.Resident, error)
	FindAll() ([]models.Resident, error)
}
type residentRepository struct {
	collection *mongo.Collection
}

func NewResidentRepository() ResidentRepository {
	return &residentRepository{
		collection: config.DB.Collection("users"),
	}
}

func (r residentRepository) Save(resident models.Resident) (models.Resident, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, resident)
	if err != nil {
		return models.Resident{}, err
	}

	resident.ID = result.InsertedID.(primitive.ObjectID)

	return resident, nil
}

func (r residentRepository) FindById(residentId string) (models.Resident, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(residentId)
	if err != nil {
		return models.Resident{}, err
	}

	var resident models.Resident

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&resident)
	if err != nil {
		return models.Resident{}, err
	}

	return resident, nil
}

func (r residentRepository) FindByEmail(email string) (models.Resident, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var resident models.Resident
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&resident)
	if err != nil {
		return models.Resident{}, err
	}

	return resident, nil
}

func (r residentRepository) FindAll() ([]models.Resident, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var residents []models.Resident
	if err := cursor.All(ctx, &residents); err != nil {
		return nil, err
	}
	return residents, nil
}

func (r residentRepository) DeleteById(residentId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert string → ObjectID
	objID, err := primitive.ObjectIDFromHex(residentId)
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
