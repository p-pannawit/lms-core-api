package model

import (
	"github.com/google/uuid"
)

type Transaction struct {
	UUID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string
	Amount   int
	Metadata []byte `gorm:"type:jsonb"`
}
