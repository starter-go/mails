package main

import (
	"os"

	"github.com/starter-go/mails/modules/mails"
	"github.com/starter-go/units"
)

func main() {

	a := os.Args
	m := mails.TestModule()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
