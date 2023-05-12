package auditlog

import (
	"time"

	"github.com/rs/xid"
)

type AuditLogSnapshot struct {
	ID           xid.ID      `json:"id" bson:"_id"`
	VariableName string      `json:"variable_name" bson:"variable_name"`
	OldValue     interface{} `json:"old_value" bson:"old_value"`
	NewValue     interface{} `json:"new_value" bson:"new_value"`
	CreatedAt    time.Time   `json:"created_at" bson:"created_at"`
}

func CreateAuditLogSnapShot(variableName string, oldValue, newValue interface{}) AuditLogSnapshot {
	als := AuditLogSnapshot{
		ID:           xid.New(),
		VariableName: variableName,
		OldValue:     oldValue,
		NewValue:     newValue,
		CreatedAt:    time.Now(),
	}

	return als
}
