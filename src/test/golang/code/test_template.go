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

// ListRegistrations implements units.Unit.
func (inst *TestTemplate) ListRegistrations(list []*units.Registration) []*units.Registration {
	list = append(list, &units.Registration{
		Name:    "test-template",
		Enabled: true,
		Do:      inst.run,
	})
	return list
}

func (inst *TestTemplate) _impl() units.Unit {
	return inst
}

func (inst *TestTemplate) run(cc context.Context) error {

	name := "demo1"
	langs := []i18n.Language{"fr_fr", "zh_cn", "en_us"} // ("zh_cn")
	tmp, err := inst.TM.Find(name, langs...)
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
