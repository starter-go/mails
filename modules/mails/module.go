package mails

import (
	"github.com/starter-go/application"
	"github.com/starter-go/mails"
	"github.com/starter-go/mails/gen/lib4mails"
	"github.com/starter-go/mails/gen/main4mails"
	"github.com/starter-go/mails/gen/test4mails"
	"github.com/starter-go/starter"
)

// LibModule ...
func LibModule() application.Module {
	mb := mails.BuildLibModule()
	mb.Components(lib4mails.ExportComponents)
	mb.Depend(starter.Module())
	return mb.Create()
}

// MainModule ...
func MainModule() application.Module {
	mb := mails.BuildMainModule()
	mb.Components(main4mails.ExportComponents)
	mb.Depend(LibModule())
	return mb.Create()
}

// TestModule ...
func TestModule() application.Module {
	mb := mails.BuildTestModule()
	mb.Components(test4mails.ExportComponents)
	mb.Depend(LibModule())
	return mb.Create()
}
