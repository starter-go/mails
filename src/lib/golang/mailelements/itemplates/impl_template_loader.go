package itemplates

import (
	"github.com/starter-go/application"
	"github.com/starter-go/i18n"
	"github.com/starter-go/mails"
	"github.com/starter-go/mails/templates"
)

// TemplateLoaderImpl ...
type TemplateLoaderImpl struct {

	//starter:component

	_as func(templates.TemplateLoader) //starter:as("#")

	I18        i18n.Service        //starter:inject("#")
	AppContext application.Context //starter:inject("context")

}

func (inst *TemplateLoaderImpl) _impl() templates.TemplateLoader { return inst }

// Load ...
func (inst *TemplateLoaderImpl) Load(name string, langs ...i18n.Language) (templates.Template, error) {

	props := inst.AppContext.GetProperties()
	g := props.Getter().Required()
	prefix := "mails.template." + name

	title := g.GetString(prefix + ".title")
	from := g.GetString(prefix + ".from-addr")
	contentType := g.GetString(prefix + ".content-type")
	contentSrc := g.GetString(prefix + ".content-source")
	defaultLang := g.GetString(prefix + ".default-language")

	err := g.Error()
	if err != nil {
		return nil, err
	}

	langs = inst.prepareLanguage(defaultLang, langs)
	res := inst.I18.GetResources(langs...)
	contentData, err := res.ReadBinary(contentSrc)
	if err != nil {
		return nil, err
	}

	msg := &mails.Message{
		Title:       title,
		FromAddress: mails.Address(from),
		Language:    langs[0],
		ContentType: contentType,
		Content:     contentData,
	}
	temp := &myTemplate{}
	temp.Message = *msg
	return temp, nil
}

func (inst *TemplateLoaderImpl) prepareLanguage(defaultLang string, langs []i18n.Language) []i18n.Language {
	size := len(langs)
	if size > 0 {
		return langs
	}
	if defaultLang == "" {
		defaultLang = "zh_cn"
	}
	langs = append(langs, i18n.Language(defaultLang))
	return langs
}
