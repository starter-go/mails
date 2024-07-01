package main4mails
import (
    pcd60e24ed "github.com/starter-go/mails/src/main/golang/code"
     "github.com/starter-go/application"
)

// type pcd60e24ed.Example in package:github.com/starter-go/mails/src/main/golang/code
//
// id:com-cd60e24ed01825b7-code-Example
// class:
// alias:
// scope:singleton
//
type pcd60e24ed0_code_Example struct {
}

func (inst* pcd60e24ed0_code_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-cd60e24ed01825b7-code-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pcd60e24ed0_code_Example) new() any {
    return &pcd60e24ed.Example{}
}

func (inst* pcd60e24ed0_code_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pcd60e24ed.Example)
	nop(ie, com)

	


    return nil
}


