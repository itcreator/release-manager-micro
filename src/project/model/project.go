package model

import "github.com/satori/go.uuid"

// Project entity store project details
type Project struct {
	//TODO: implement DDL "id UUID PRIMARY KEY DEFAULT gen_random_uuid(),"
	UUID        uuid.UUID `db:"uuid"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}
