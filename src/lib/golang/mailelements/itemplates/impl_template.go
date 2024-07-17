package itemplates

import (
	"strings"
	"text/template"

	"github.com/starter-go/mails"
	"github.com/starter-go/mails/templates"
)

type myTemplate struct {
	mails.Message
}

func (inst *myTemplate) _impl() templates.Template { return inst }

func (inst *myTemplate) NewMessage(props map[string]string) (*mails.Message, error) {

	src := &inst.Message
	dst := &mails.Message{}
	*dst = *src

	text1 := string(src.Content)
	text2 := inst.makeTextContent(text1, props)

	dst.Content = []byte(text2)
	return dst, nil
}

func (inst *myTemplate) makeTextContent(text string, props map[string]string) string {

	buffer := &strings.Builder{}
	t, err := template.New("mails.message.template").Parse(text)
	if err != nil {
		return "err:" + err.Error()
	}

	err = t.Execute(buffer, props)
	if err != nil {
		return "err:" + err.Error()
	}

	return buffer.String()
}
