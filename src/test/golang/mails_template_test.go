package main

import (
	"testing"

	"github.com/starter-go/mails/modules/mails"
	"github.com/starter-go/units"
)

func runTestCase(t *testing.T, caseName string) {

	args := []string{}
	mod := mails.TestModule()

	units.Run(&units.Config{
		T:        t,
		Args:     args,
		Cases:    caseName,
		Module:   mod,
		UsePanic: true,
	})

}

////////////////////////////////////////////////////////////////////////////////

func TestMailsTemplate(t *testing.T) {
	runTestCase(t, "test-template")
}

func Test1(t *testing.T) {
	runTestCase(t, "test-1")
}
