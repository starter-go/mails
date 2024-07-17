package test4mails
import (
    pd671d76a1 "github.com/starter-go/mails"
    p405cc385e "github.com/starter-go/mails/src/test/golang/code"
    p69cb21798 "github.com/starter-go/mails/templates"
     "github.com/starter-go/application"
)

// type p405cc385e.Test1 in package:github.com/starter-go/mails/src/test/golang/code
//
// id:com-405cc385eb732ce5-code-Test1
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p405cc385eb_code_Test1 struct {
}

func (inst* p405cc385eb_code_Test1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-405cc385eb732ce5-code-Test1"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
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



// type p405cc385e.TestTemplate in package:github.com/starter-go/mails/src/test/golang/code
//
// id:com-405cc385eb732ce5-code-TestTemplate
// class:
// alias:
// scope:singleton
//
type p405cc385eb_code_TestTemplate struct {
}

func (inst* p405cc385eb_code_TestTemplate) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-405cc385eb732ce5-code-TestTemplate"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p405cc385eb_code_TestTemplate) new() any {
    return &p405cc385e.TestTemplate{}
}

func (inst* p405cc385eb_code_TestTemplate) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p405cc385e.TestTemplate)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.TM = inst.getTM(ie)


    return nil
}


func (inst*p405cc385eb_code_TestTemplate) getSender(ie application.InjectionExt)pd671d76a1.Service{
    return ie.GetComponent("#alias-d671d76a169061f84f6814f84b98af24-Service").(pd671d76a1.Service)
}


func (inst*p405cc385eb_code_TestTemplate) getTM(ie application.InjectionExt)p69cb21798.TemplateManager{
    return ie.GetComponent("#alias-69cb21798ba841067147ba8fee303bf2-TemplateManager").(p69cb21798.TemplateManager)
}


