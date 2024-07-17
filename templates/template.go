package templates

import (
	"github.com/starter-go/i18n"
	"github.com/starter-go/mails"
)

// Template 代表一个模板
type Template interface {
	NewMessage(props map[string]string) (*mails.Message, error)
}

// TemplateManager 是模板管理器
type TemplateManager interface {

	// 查找名为 'name' 的模板
	Find(name string, langs ...i18n.Language) (Template, error)
}

// TemplateCache 缓存已加载的模板
type TemplateCache interface {
	Fetch(name string, langs ...i18n.Language) (Template, error)
}

// TemplateLoader 用于加载模板
type TemplateLoader interface {

	//  加载名为 'name' 的模板
	Load(name string, langs ...i18n.Language) (Template, error)
}
