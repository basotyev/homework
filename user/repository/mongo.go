package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"lesson13/models"
	"lesson13/user"
)

type MongoRepository struct {
	db *mongo.Client
}

func NewMongoRepository(db *mongo.Client) user.Repository {
	return &MongoRepository{
		db: db,
	}
}

func (m *MongoRepository) CreateUser(ctx context.Context, user *models.User) error {
	res, err := m.db.Database("test-db").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	user.MongoId = res.InsertedID.(primitive.ObjectID).String()
	return nil
}

func (m *MongoRepository) UpdateUserById(ctx context.Context, user *models.User) error {
	_, err := m.db.Database("test-db").Collection("users").UpdateOne(ctx, bson.M{"id": user.MongoId}, user)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoRepository) GetUserById(ctx context.Context, id int) (*models.User, error) {
	res, err := m.db.Database("test-db").Collection("users").Find(ctx, bson.M{"id": id})
	if err != nil {
		return nil, err
	}
	var usr models.User
	err = res.Decode(&usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (m *MongoRepository) RemoveUserById(ctx context.Context, id int) error {
	_, err := m.db.Database("test-db").Collection("users").DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}
