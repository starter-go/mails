package code

import (
	"context"

	"github.com/starter-go/i18n"
	"github.com/starter-go/mails"
	"github.com/starter-go/mails/templates"
	"github.com/starter-go/units"
)

// TestTemplate ...
type TestTemplate struct {

	//starter:component

	// _as func(units.Units) //starter:as(".")

	Sender mails.Service             //starter:inject("#")
	TM     templates.TemplateManager //starter:inject("#")

	// ToAddr string //starter:inject("${mails.test.to-addr}")
}

func (inst *TestTemplate) _impl() units.Units { return inst }

// Units ...
func (inst *TestTemplate) Units(list []*units.Registration) []*units.Registration {
	list = append(list, &units.Registration{
		Name:    "test-template",
		Enabled: true,
		Test:    inst.run,
	})
	return list
}

func (inst *TestTemplate) run() error {

	name := "demo1"
	lang := i18n.Language("zh_cn")
	tmp, err := inst.TM.Find(name, lang)
	if err != nil {
		return err
	}

	props := map[string]string{
		"a": "apple1",
		"b": "beats2",
	}
	msg, err := tmp.NewMessage(props)

	ctx := context.Background()
	return inst.Sender.Send(ctx, msg)
}

// func (inst *TestTemplate) loop() error {
// 	for {
// 		time.Sleep(time.Second)
// 	}
// }
