package auditlog

import (
	"context"
	"fmt"
)

type AuditLogService interface {
	CreateSnapshot(ctx context.Context, variableName string, oldValue, newValue interface{}) error
}

type service struct {
	storage Storage
}

func NewAuditLog(storage Storage) AuditLogService {
	return &service{
		storage: storage,
	}
}

func (s *service) CreateSnapshot(ctx context.Context, variableName string, oldValue, newValue interface{}) error {
	snapshot := CreateAuditLogSnapShot(variableName, oldValue, newValue)

	fmt.Printf("Variable: %s changed from: %+v to:%+v\n", variableName, oldValue, newValue)

	_, err := s.storage.Insert(ctx, snapshot)
	if err != nil {
		return fmt.Errorf("failed to insert snapshot: %w", err)
	}

	return nil
}
