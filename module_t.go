package mails

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName     = "github.com/starter-go/mails"
	theModuleVersion  = "v0.0.5"
	theModuleRevision = 5
)

////////////////////////////////////////////////////////////////////////////////

const (
	theLibModuleResPath  = "src/lib/resources"
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/lib/resources"
var theLibModuleResFS embed.FS

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// BuildLibModule 导出模块
func BuildLibModule() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#lib")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theLibModuleResFS, theLibModuleResPath)
	return mb
}

// BuildMainModule 导出模块
func BuildMainModule() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	return mb
}

// BuildTestModule 导出模块
func BuildTestModule() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}
