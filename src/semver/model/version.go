package model

// Version entity for semantic versioning strategy
type Version struct {
	ID        uint64 `db:"id"`
	ProjectID uint64 `db:"project_id"`
	Major     uint32 `db:"major"`
	Minor     uint32 `db:"minor"`
	Revision  uint32 `db:"revision"`
	Branch    string `db:"branch"`
}
