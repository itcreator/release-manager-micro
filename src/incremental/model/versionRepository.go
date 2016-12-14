package model

import (
	"database/sql"
	"gopkg.in/gorp.v1"
)

var tableName = "version_incremental"

// IVersionRepository is interface for version repository
type IVersionRepository interface {
	Insert(ver *Version) *Version
	Select(projectName string) (*Version, bool)
	Update(ver *Version) *Version
	Delete(projectName string)
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
func (rep *VersionRepository) Select(projectName string) (*Version, bool) {
	v := new(Version)
	query := "SELECT * FROM " + tableName
	query += " WHERE project_name = :projectName;"

	params := map[string]interface{}{
		"projectName": projectName,
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

// Update function update revision
func (rep *VersionRepository) Update(ver *Version) *Version {
	stmt, err := rep.DbMap.Prepare("UPDATE " + tableName + " SET version=$1 WHERE project_name=$2")
	checkErr(err)

	res, err := stmt.Exec(ver.Revision, ver.ProjectName)
	checkErr(err)
	res.RowsAffected()

	return ver
}

// Delete function delete version from DB
func (rep *VersionRepository) Delete(projectName string) {
	stmt, err := rep.DbMap.Prepare("DELETE FROM " + tableName + " WHERE project_name=$1")
	checkErr(err)

	res, err := stmt.Exec(projectName)
	checkErr(err)
	res.RowsAffected()
}
