package main

import (
	"os"

	"github.com/starter-go/mails/modules/mails"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(mails.LibModule())
	i.WithPanic(true).Run()
}
