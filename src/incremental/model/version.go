package model

// Version entity for incremental versioning strategy
type Version struct {
	ID          uint64 `db:"id"`
	ProjectName string `db:"project_name"`
	Revision    uint64 `db:"version"`
}
