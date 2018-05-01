package model

// Set of tags for generated version
type TagSet struct {
	IsLatest bool
	Full     string
	Major    *string
	Minor    *string
	Branch   *string
}
