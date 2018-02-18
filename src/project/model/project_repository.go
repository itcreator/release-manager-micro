package model

import (
	"fmt"
	"gopkg.in/gorp.v2"
)

var tableName = "project"

//IProjectRepository interface for project DB storage
type IProjectRepository interface {
	Insert(p *Project)
	SelectByID(id uint64) *Project
	Update(p *Project) bool
	SelectAll() []*Project
}

// ProjectRepository is table rep for `Project`
type ProjectRepository struct {
	DbMap *gorp.DbMap `inject:"dbMap"`
}

//Insert new version into db
func (rep *ProjectRepository) Insert(p *Project) {
	err := rep.DbMap.Insert(p)
	checkErr(err)
}

// SelectByID selects project from DB
func (rep *ProjectRepository) SelectByID(id uint64) *Project {
	var p *Project
	obj, err := rep.DbMap.Get(Project{}, id)
	checkErr(err)
	if nil != obj {
		p = obj.(*Project)
	}

	return p
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
}

//Update one project in DB
func (rep *ProjectRepository) Update(p *Project) bool {
	count, err := rep.DbMap.Update(p)
	checkErr(err)

	return count > 0
}

// SelectAll select all projects from DB
func (rep *ProjectRepository) SelectAll() []*Project {
	var projects []*Project
	_, err := rep.DbMap.Select(&projects, "SELECT * FROM "+tableName+";")

	checkErr(err)

	return projects
}
