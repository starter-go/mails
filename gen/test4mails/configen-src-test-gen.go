package test4mails
import (
    pd671d76a1 "github.com/starter-go/mails"
    p405cc385e "github.com/starter-go/mails/src/test/golang/code"
     "github.com/starter-go/application"
)

// type p405cc385e.Test1 in package:github.com/starter-go/mails/src/test/golang/code
//
// id:com-405cc385eb732ce5-code-Test1
// class:
// alias:
// scope:singleton
//
type p405cc385eb_code_Test1 struct {
}

func (inst* p405cc385eb_code_Test1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-405cc385eb732ce5-code-Test1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p405cc385eb_code_Test1) new() any {
    return &p405cc385e.Test1{}
}

func (inst* p405cc385eb_code_Test1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p405cc385e.Test1)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.ToAddr = inst.getToAddr(ie)


    return nil
}


func (inst*p405cc385eb_code_Test1) getSender(ie application.InjectionExt)pd671d76a1.Service{
    return ie.GetComponent("#alias-d671d76a169061f84f6814f84b98af24-Service").(pd671d76a1.Service)
}


func (inst*p405cc385eb_code_Test1) getToAddr(ie application.InjectionExt)string{
    return ie.GetString("${mails.test.to-addr}")
}


