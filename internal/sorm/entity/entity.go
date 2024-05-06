package entity

import (
	"encoding/json"
)

type IEntity interface {
	GetRules() ValidationRules
	GetCorrelationId() string
	GetName() string
}

func ToBytes(entity IEntity) []byte {
	bites, _ := json.Marshal(entity)
	return bites
}
