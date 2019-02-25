package model

import (
	"github.com/satori/go.uuid"
)

// Version entity for semantic versioning strategy
type Version struct {
	UUID        uuid.NullUUID `db:"uuid"`
	ProjectUUID uuid.UUID     `db:"project_uuid"`
	Major       uint32        `db:"major"`
	Minor       uint32        `db:"minor"`
	Revision    uint32        `db:"revision"`
	Branch      string        `db:"branch"`
}
