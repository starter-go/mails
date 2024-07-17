package itemplates

import (
	"strings"
	"sync"

	"github.com/starter-go/i18n"
	"github.com/starter-go/mails/templates"
)

// TemplateCacheImpl ...
type TemplateCacheImpl struct {

	//starter:component

	_as func(templates.TemplateCache) //starter:as("#")

	Next templates.TemplateLoader //starter:inject("#")

	table map[string]templates.Template
	mutex sync.Mutex
}

func (inst *TemplateCacheImpl) _impl() templates.TemplateCache { return inst }

// Fetch ...
func (inst *TemplateCacheImpl) Fetch(name string, langs ...i18n.Language) (templates.Template, error) {

	mu := &inst.mutex
	mu.Lock()
	defer mu.Unlock()

	// find in cache
	key := inst.keyFor(name, langs...)
	table := inst.getTable()
	temp := table[key]
	if temp != nil {
		return temp, nil
	}

	// load
	t2, err := inst.Next.Load(name, langs...)
	if err != nil {
		return nil, err
	}
	temp = t2

	// cache
	table[key] = temp
	return temp, nil
}

func (inst *TemplateCacheImpl) keyFor(name string, langs ...i18n.Language) string {
	builder := &strings.Builder{}
	builder.WriteString(name)
	builder.WriteString(".")
	for _, item := range langs {
		builder.WriteString(item.String())
		break
	}
	builder.WriteString(".t")
	str := builder.String()
	return strings.ToLower(str)
}

func (inst *TemplateCacheImpl) getTable() map[string]templates.Template {
	t := inst.table
	if t == nil {
		t = make(map[string]templates.Template)
		inst.table = t
	}
	return t
}
