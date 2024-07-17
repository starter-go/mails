package mails

import (
	"github.com/starter-go/application"
	"github.com/starter-go/i18n/modules/i18n"
	"github.com/starter-go/mails"
	"github.com/starter-go/mails/gen/lib4mails"
	"github.com/starter-go/mails/gen/main4mails"
	"github.com/starter-go/mails/gen/test4mails"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
)

// LibModule ...
func LibModule() application.Module {
	mb := mails.BuildLibModule()
	mb.Components(lib4mails.ExportComponents)
	mb.Depend(starter.Module())
	mb.Depend(i18n.Module())
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
	mb.Depend(units.Module())
	return mb.Create()
}
