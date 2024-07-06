// infrastructure/database/user_repository.go
package database

import (
	"BCScanner/domain"
	"BCScanner/interfaces"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) interfaces.UserRepository {
	return &UserRepository{Collection: db.Collection("users")}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), user)
	return err
}

func (r *UserRepository) GetUserByID(id string) (*domain.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user domain.User
	err = r.Collection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&user)
	return &user, err
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
		options.Update().SetUpsert(false),
	)
	return err
}

func (r *UserRepository) DeleteUser(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": oid})
	return err
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	cursor, err := r.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var users []domain.User
	err = cursor.All(context.Background(), &users)
	return users, err
}
