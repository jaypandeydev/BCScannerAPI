package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionQRImport = "QrData"
)

type QRGenerateRequest struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	QRImageURL string             `bson:"qrImageUrl" json:"qrImageUrl"`
	OwnerName  string             `bson:"ownerName" json:"ownerName"`
	CreatedAt  primitive.DateTime `bson:"createdAt" json:"createdAt"`
	Status     bool               `bson:"status" json:"status"`
}

type QRGenerateResponse struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	QRImageURL string             `bson:"qrImageUrl" json:"qrImageUrl"`
	OwnerName  string             `bson:"ownerName" json:"ownerName"`
	CreatedAt  primitive.DateTime `bson:"createdAt" json:"createdAt"`
	Status     bool               `bson:"status" json:"status"`
}

type QRGenerateUseCase interface {
	CreateQR(ctx context.Context, qrGenerator *QRGenerateRequest) (InsertedID string, err error)
	Fetch(ctx context.Context) ([]bson.D, error)
	GetByID(ctx context.Context, id string) (QRGenerateResponse, error)
}
type QRGenerateRepository interface {
	CreateQR(ctx context.Context, qrGenerator *QRGenerateRequest) (InsertedID string, err error)
	Fetch(ctx context.Context) ([]bson.D, error)
	GetByID(ctx context.Context, id string) (QRGenerateResponse, error)
}
