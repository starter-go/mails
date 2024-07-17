package code

import (
	"context"

	"github.com/starter-go/mails"
	"github.com/starter-go/units"
)

// Test1 ...
type Test1 struct {

	//starter:component

	_as func(units.Units) //starter:as(".")

	Sender mails.Service //starter:inject("#")

	ToAddr string //starter:inject("${mails.test.to-addr}")
}

func (inst *Test1) _impl() units.Units { return inst }

// Units ...
func (inst *Test1) Units(list []*units.Registration) []*units.Registration {
	list = append(list, &units.Registration{
		Name:    "test-1",
		Enabled: true,
		Test:    inst.run,
	})
	return list
}

func (inst *Test1) run() error {

	ctx := context.Background()
	msg := &mails.Message{}

	text := "hello, world"
	toAddr := inst.ToAddr

	msg.ToAddresses = []mails.Address{mails.Address(toAddr)}
	msg.Title = "hello"
	msg.ContentType = "text/plain"
	msg.Content = []byte(text)

	return inst.Sender.Send(ctx, msg)
}

// func (inst *Test1) loop() error {
// 	for {
// 		time.Sleep(time.Second)
// 	}
// }
