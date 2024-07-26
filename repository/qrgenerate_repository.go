package repository

import (
	"BCScanner/domain"
	"BCScanner/infrastructure/mongo"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	//"log"
)

type QRGenerateRepository struct {
	database   mongo.Database
	collection string
}

func NewQRGenerateRepository(database mongo.Database, collection string) domain.QRGenerateRepository {
	return &QRGenerateRepository{
		database:   database,
		collection: collection,
	}
}

func (qr *QRGenerateRepository) CreateQR(ctx context.Context, qrGeneratorReq *domain.QRGenerateRequest) (string, error) {
	collection := qr.database.Collection(qr.collection)

	var InsertedID string
	// Insert the document
	results, err := collection.InsertOne1(ctx, qrGeneratorReq)
	if err != nil {
		return InsertedID, err
	}
	id, ok := results.InsertedID.(primitive.ObjectID)
	if !ok {
		return InsertedID, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}
	InsertedID = id.Hex()
	return InsertedID, nil
}

func (qr *QRGenerateRepository) Fetch(ctx context.Context) ([]bson.D, error) {
	collection := qr.database.Collection(qr.collection)

	// Define options (optional)
	findOptions := options.Find()
	findOptions.SetLimit(10) // Limit to 10 documents

	// Find documents
	cur, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	// Defer the cursor close and handle any error
	defer func() {
		if closeErr := cur.Close(ctx); closeErr != nil {
			log.Printf("Error closing cursor: %v", closeErr)
		}
	}()

	// Iterate through the cursor and decode each document
	var results []bson.D
	for cur.Next(ctx) {
		var result bson.D
		if err := cur.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	// Check for any cursor errors
	if err := cur.Close(ctx); err != nil {
		return nil, err
	}

	return results, nil
}

func (qr *QRGenerateRepository) GetByID(ctx context.Context, id string) (domain.QRGenerateResponse, error) {
	collection := qr.database.Collection(qr.collection)

	var QRGenerate domain.QRGenerateResponse

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return QRGenerate, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(&QRGenerate)
	return QRGenerate, err
}
