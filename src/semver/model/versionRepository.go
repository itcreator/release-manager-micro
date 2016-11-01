package model

import (
	"database/sql"
	"fmt"
	"gopkg.in/gorp.v1"
)

var tableName string = "version_semantic"

// IVersionRepository is interface for version repository
type IVersionRepository interface {
	Insert(ver *Version) *Version
	Select(projectId uint64, major uint32, minor uint32, branch string) (*Version, bool)
	UpdateRevision(ver *Version) *Version
}

// VersionRepository is data repository for `Version`
type VersionRepository struct {
	DbMap *gorp.DbMap `inject:"dbMap"`
}

//Insert new version
func (rep *VersionRepository) Insert(ver *Version) *Version {
	// insert
	fields := "(project_id, major, minor, revision, branch)"
	values := "($1, $2, $3, $4, $5)"
	stmt, err := rep.DbMap.Prepare("INSERT INTO " + tableName + fields + "VALUES" + values + " returning id;")
	checkErr(err)

	//last insert id for postgres
	row := stmt.QueryRow(ver.ProjectId, ver.Major, ver.Minor, ver.Revision, ver.Branch)

	err = row.Scan(&ver.Id)
	if err != sql.ErrNoRows {
		checkErr(err)
	}

	fmt.Println(ver.Id)

	return ver
}

// Select version from DB
func (rep *VersionRepository) Select(projectId uint64, major uint32, minor uint32, branch string) (*Version, bool) {
	stmt, err := rep.DbMap.Prepare("SELECT id, project_id, major, minor, revision, branch FROM " + tableName + " WHERE project_id = $1 AND major = $2 AND minor = $3 and branch= $4;")
	checkErr(err)

	v := new(Version)
	row := stmt.QueryRow(projectId, major, minor, branch) //id=1
	err = row.Scan(&v.Id, &v.ProjectId, &v.Major, &v.Minor, &v.Revision, &v.Branch)
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

	res, err := stmt.Exec(ver.Revision, ver.Id)
	checkErr(err)
	res.RowsAffected()

	return ver
}
