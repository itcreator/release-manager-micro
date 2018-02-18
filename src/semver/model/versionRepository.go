package model

import (
	"database/sql"
	"gopkg.in/gorp.v2"
)

var tableName = "version_semantic"

// IVersionRepository is interface for version repository
type IVersionRepository interface {
	Insert(ver *Version) *Version
	Select(projectID uint64, major uint32, minor uint32, branch string) (*Version, bool)
	UpdateRevision(ver *Version) *Version
}

// VersionRepository is data repository for `Version`
type VersionRepository struct {
	DbMap *gorp.DbMap `inject:"dbMap"`
}

//Insert new version
func (rep *VersionRepository) Insert(ver *Version) *Version {
	err := rep.DbMap.Insert(ver)
	checkErr(err)

	return ver
}

// Select version from DB
func (rep *VersionRepository) Select(projectID uint64, major uint32, minor uint32, branch string) (*Version, bool) {
	v := new(Version)
	query := "SELECT id, project_id, major, minor, revision, branch FROM " + tableName
	query += " WHERE project_id = :projectId AND major = :major AND minor = :minor AND branch = :branch;"

	params := map[string]interface{}{
		"projectId": projectID,
		"major":     major,
		"minor":     minor,
		"branch":    branch,
	}
	err := rep.DbMap.SelectOne(v, query, params)
	if err != sql.ErrNoRows {
		checkErr(err)
	}

	isEmpty := err == sql.ErrNoRows

	return v, isEmpty
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// UpdateRevision function update only revision field
func (rep *VersionRepository) UpdateRevision(ver *Version) *Version {
	stmt, err := rep.DbMap.Prepare("UPDATE " + tableName + " SET revision=$1 WHERE id=$2")

	res, err := stmt.Exec(ver.Revision, ver.ID)
	checkErr(err)
	res.RowsAffected()

	return ver
}
