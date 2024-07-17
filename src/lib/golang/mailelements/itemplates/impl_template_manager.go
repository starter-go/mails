package itemplates

import (
	"github.com/starter-go/i18n"
	"github.com/starter-go/mails/templates"
)

// TemplateManagerImpl ...
type TemplateManagerImpl struct {

	//starter:component

	_as func(templates.TemplateManager) //starter:as("#")

	Next templates.TemplateCache //starter:inject("#")

}

func (inst *TemplateManagerImpl) _impl() templates.TemplateManager { return inst }

// Find ...
func (inst *TemplateManagerImpl) Find(name string, langs ...i18n.Language) (templates.Template, error) {
	return inst.Next.Fetch(name, langs...)
}
