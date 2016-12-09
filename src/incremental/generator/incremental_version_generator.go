package generator

import (
	"incremental/model"
)

// IIncrementalVersionGenerator generate version
type IIncrementalVersionGenerator interface {
	GenerateVersion(projectName string) uint64
	DeleteVersion(projectName string)
}

// IncrementalVersionGenerator implements `IIncrementalVersionGenerator` interface
type IncrementalVersionGenerator struct {
	VersionRepository model.IVersionRepository `inject:""`
}

func (g *IncrementalVersionGenerator) getStoredVersion(projectName string) (*model.Version, bool) {

	rep := g.VersionRepository
	ver, isEmpty := rep.Select(projectName)

	if isEmpty {
		ver.ProjectName = projectName
		ver.Revision = 1
	} else {
		ver.Revision++
	}

	return ver, isEmpty
}

// GenerateVersion function generate version for project
func (g *IncrementalVersionGenerator) GenerateVersion(projectName string) uint64 {

	ver, isNew := g.getStoredVersion(projectName)

	rep := g.VersionRepository
	if isNew {
		rep.Insert(ver)
	} else {
		rep.Update(ver)
	}

	return ver.Revision
}

// DeleteVersion function delete generated version
func (g *IncrementalVersionGenerator) DeleteVersion(projectName string) {
	g.VersionRepository.Delete(projectName)
}
