package model

// Project entity store project details
type Project struct {
	ID          uint64 `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
