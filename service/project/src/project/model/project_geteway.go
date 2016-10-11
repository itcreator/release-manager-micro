package model

import (
	"database/sql"
	"fmt"
)

var tableName string = "project"

//IProjectGateway interface for project DB storage
type IProjectGateway interface {
	Insert(p *Project)
	SelectById(id uint64) (*Project, bool)
	Update(p *Project) *Project
	SelectAll() []*Project
}

// ProjectGateway is table gateway for `Project`
type ProjectGateway struct {
	Db *sql.DB `inject:"db"`
}

//Insert new version into db
func (gateway *ProjectGateway) Insert(p *Project) {
	// insert
	fields := " (name, description) "
	values := "($1, $2)"
	stmt, err := gateway.Db.Prepare("INSERT INTO " + tableName + fields + "VALUES" + values + " returning id;")
	checkErr(err)
	defer stmt.Close()

	//last insert id for postgres
	row := stmt.QueryRow(p.Name, p.Description)

	err = row.Scan(&p.Id)
	if err != sql.ErrNoRows {
		checkErr(err)
	}
}

// SelectById selects project from DB
func (gateway *ProjectGateway) SelectById(id uint64) (*Project, bool) {
	stmt, err := gateway.Db.Prepare("SELECT id, name, description FROM " + tableName + " WHERE id = $1;")
	checkErr(err)
	defer stmt.Close()

	p := new(Project)
	row := stmt.QueryRow(id) //id=1
	err = row.Scan(&p.Id, &p.Name, &p.Description)
	if err != sql.ErrNoRows {
		checkErr(err)
	}

	isNotFound := err == sql.ErrNoRows

	return p, isNotFound
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

//Update one project in DB
func (gateway *ProjectGateway) Update(p *Project) *Project {
	stmt, err := gateway.Db.Prepare("UPDATE " + tableName + " SET name=$1, description=$2 WHERE id=$3")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Description, p.Id)
	checkErr(err)

	return p
}

// SelectAll select all projects from DB
func (gateway *ProjectGateway) SelectAll() []*Project {
	stmt, err := gateway.Db.Prepare("SELECT id, name, description FROM " + tableName + ";")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil && err != sql.ErrNoRows {
		checkErr(err)
	}
	defer rows.Close()

	var projects = []*Project{}
	for rows.Next() {
		p := new(Project)
		err := rows.Scan(&p.Id, &p.Name, &p.Description)
		if err != nil {
			checkErr(err)
		}
		projects = append(projects, p)
	}

	err = rows.Err()
	if err != nil {
		checkErr(err)
	}

	return projects
}
