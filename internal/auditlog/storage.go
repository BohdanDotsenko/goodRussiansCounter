package auditlog

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage interface {
	Insert(ctx context.Context, p AuditLogSnapshot) (AuditLogSnapshot, error)
}

type StorageMongo struct {
	collection *mongo.Collection
}

func NewStorageMongo(
	mongoCollection *mongo.Collection,
) *StorageMongo {
	return &StorageMongo{
		collection: mongoCollection,
	}
}

func (s *StorageMongo) Insert(ctx context.Context, p AuditLogSnapshot) (AuditLogSnapshot, error) {
	_, err := s.collection.InsertOne(ctx, p)
	if err != nil {
		return AuditLogSnapshot{}, err
	}

	return p, nil
}
