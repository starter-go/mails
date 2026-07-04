package main

import (
	"testing"

	"github.com/starter-go/mails/modules/mails"
	"github.com/starter-go/units"
)

func runTestCase(t *testing.T, caseName string) {

	a := []string{}
	m := mails.TestModule()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
		T:         t,
		Selector:  caseName,
	}

	units.Run(c)
}

////////////////////////////////////////////////////////////////////////////////

func TestMailsTemplate(t *testing.T) {
	runTestCase(t, "test-template")
}

func Test1(t *testing.T) {
	runTestCase(t, "test-1")
}
