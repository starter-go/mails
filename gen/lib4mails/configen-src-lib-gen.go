package lib4mails
import (
    p0ef6f2938 "github.com/starter-go/application"
    p52d293b71 "github.com/starter-go/i18n"
    pd671d76a1 "github.com/starter-go/mails"
    p1f13f945f "github.com/starter-go/mails/src/lib/golang/mailelements/implements"
    p90e018b46 "github.com/starter-go/mails/src/lib/golang/mailelements/itemplates"
    p69cb21798 "github.com/starter-go/mails/templates"
     "github.com/starter-go/application"
)

// type p1f13f945f.MainDispatcherRegistry in package:github.com/starter-go/mails/src/lib/golang/mailelements/implements
//
// id:com-1f13f945f51fcda2-implements-MainDispatcherRegistry
// class:class-d671d76a169061f84f6814f84b98af24-DispatcherRegistry
// alias:
// scope:singleton
//
type p1f13f945f5_implements_MainDispatcherRegistry struct {
}

func (inst* p1f13f945f5_implements_MainDispatcherRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1f13f945f51fcda2-implements-MainDispatcherRegistry"
	r.Classes = "class-d671d76a169061f84f6814f84b98af24-DispatcherRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1f13f945f5_implements_MainDispatcherRegistry) new() any {
    return &p1f13f945f.MainDispatcherRegistry{}
}

func (inst* p1f13f945f5_implements_MainDispatcherRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1f13f945f.MainDispatcherRegistry)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)
    com.Drivers = inst.getDrivers(ie)
    com.DispatcherNameList = inst.getDispatcherNameList(ie)


    return nil
}


func (inst*p1f13f945f5_implements_MainDispatcherRegistry) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p1f13f945f5_implements_MainDispatcherRegistry) getDrivers(ie application.InjectionExt)pd671d76a1.DriverManager{
    return ie.GetComponent("#alias-d671d76a169061f84f6814f84b98af24-DriverManager").(pd671d76a1.DriverManager)
}


func (inst*p1f13f945f5_implements_MainDispatcherRegistry) getDispatcherNameList(ie application.InjectionExt)string{
    return ie.GetString("${mails.dispatcher-name-list}")
}



// type p1f13f945f.DriverManagerImpl in package:github.com/starter-go/mails/src/lib/golang/mailelements/implements
//
// id:com-1f13f945f51fcda2-implements-DriverManagerImpl
// class:
// alias:alias-d671d76a169061f84f6814f84b98af24-DriverManager
// scope:singleton
//
type p1f13f945f5_implements_DriverManagerImpl struct {
}

func (inst* p1f13f945f5_implements_DriverManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1f13f945f51fcda2-implements-DriverManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-d671d76a169061f84f6814f84b98af24-DriverManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1f13f945f5_implements_DriverManagerImpl) new() any {
    return &p1f13f945f.DriverManagerImpl{}
}

func (inst* p1f13f945f5_implements_DriverManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1f13f945f.DriverManagerImpl)
	nop(ie, com)

	
    com.Regs = inst.getRegs(ie)


    return nil
}


func (inst*p1f13f945f5_implements_DriverManagerImpl) getRegs(ie application.InjectionExt)[]pd671d76a1.DriverRegistry{
    dst := make([]pd671d76a1.DriverRegistry, 0)
    src := ie.ListComponents(".class-d671d76a169061f84f6814f84b98af24-DriverRegistry")
    for _, item1 := range src {
        item2 := item1.(pd671d76a1.DriverRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p1f13f945f.DispatcherManager in package:github.com/starter-go/mails/src/lib/golang/mailelements/implements
//
// id:com-1f13f945f51fcda2-implements-DispatcherManager
// class:
// alias:alias-d671d76a169061f84f6814f84b98af24-Service
// scope:singleton
//
type p1f13f945f5_implements_DispatcherManager struct {
}

func (inst* p1f13f945f5_implements_DispatcherManager) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1f13f945f51fcda2-implements-DispatcherManager"
	r.Classes = ""
	r.Aliases = "alias-d671d76a169061f84f6814f84b98af24-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1f13f945f5_implements_DispatcherManager) new() any {
    return &p1f13f945f.DispatcherManager{}
}

func (inst* p1f13f945f5_implements_DispatcherManager) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1f13f945f.DispatcherManager)
	nop(ie, com)

	
    com.Regs = inst.getRegs(ie)
    com.DefaultSenderAddr = inst.getDefaultSenderAddr(ie)


    return nil
}


func (inst*p1f13f945f5_implements_DispatcherManager) getRegs(ie application.InjectionExt)[]pd671d76a1.DispatcherRegistry{
    dst := make([]pd671d76a1.DispatcherRegistry, 0)
    src := ie.ListComponents(".class-d671d76a169061f84f6814f84b98af24-DispatcherRegistry")
    for _, item1 := range src {
        item2 := item1.(pd671d76a1.DispatcherRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p1f13f945f5_implements_DispatcherManager) getDefaultSenderAddr(ie application.InjectionExt)string{
    return ie.GetString("${mails.default-sender-address}")
}



// type p1f13f945f.MockDriver in package:github.com/starter-go/mails/src/lib/golang/mailelements/implements
//
// id:com-1f13f945f51fcda2-implements-MockDriver
// class:class-d671d76a169061f84f6814f84b98af24-DriverRegistry
// alias:
// scope:singleton
//
type p1f13f945f5_implements_MockDriver struct {
}

func (inst* p1f13f945f5_implements_MockDriver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1f13f945f51fcda2-implements-MockDriver"
	r.Classes = "class-d671d76a169061f84f6814f84b98af24-DriverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1f13f945f5_implements_MockDriver) new() any {
    return &p1f13f945f.MockDriver{}
}

func (inst* p1f13f945f5_implements_MockDriver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1f13f945f.MockDriver)
	nop(ie, com)

	


    return nil
}



// type p1f13f945f.SMSDriver in package:github.com/starter-go/mails/src/lib/golang/mailelements/implements
//
// id:com-1f13f945f51fcda2-implements-SMSDriver
// class:class-d671d76a169061f84f6814f84b98af24-DriverRegistry
// alias:
// scope:singleton
//
type p1f13f945f5_implements_SMSDriver struct {
}

func (inst* p1f13f945f5_implements_SMSDriver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1f13f945f51fcda2-implements-SMSDriver"
	r.Classes = "class-d671d76a169061f84f6814f84b98af24-DriverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1f13f945f5_implements_SMSDriver) new() any {
    return &p1f13f945f.SMSDriver{}
}

func (inst* p1f13f945f5_implements_SMSDriver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1f13f945f.SMSDriver)
	nop(ie, com)

	


    return nil
}



// type p1f13f945f.SMTPSenderDriver in package:github.com/starter-go/mails/src/lib/golang/mailelements/implements
//
// id:com-1f13f945f51fcda2-implements-SMTPSenderDriver
// class:class-d671d76a169061f84f6814f84b98af24-DriverRegistry
// alias:
// scope:singleton
//
type p1f13f945f5_implements_SMTPSenderDriver struct {
}

func (inst* p1f13f945f5_implements_SMTPSenderDriver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1f13f945f51fcda2-implements-SMTPSenderDriver"
	r.Classes = "class-d671d76a169061f84f6814f84b98af24-DriverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1f13f945f5_implements_SMTPSenderDriver) new() any {
    return &p1f13f945f.SMTPSenderDriver{}
}

func (inst* p1f13f945f5_implements_SMTPSenderDriver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1f13f945f.SMTPSenderDriver)
	nop(ie, com)

	


    return nil
}



// type p90e018b46.TemplateCacheImpl in package:github.com/starter-go/mails/src/lib/golang/mailelements/itemplates
//
// id:com-90e018b46a50d1bf-itemplates-TemplateCacheImpl
// class:
// alias:alias-69cb21798ba841067147ba8fee303bf2-TemplateCache
// scope:singleton
//
type p90e018b46a_itemplates_TemplateCacheImpl struct {
}

func (inst* p90e018b46a_itemplates_TemplateCacheImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-90e018b46a50d1bf-itemplates-TemplateCacheImpl"
	r.Classes = ""
	r.Aliases = "alias-69cb21798ba841067147ba8fee303bf2-TemplateCache"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p90e018b46a_itemplates_TemplateCacheImpl) new() any {
    return &p90e018b46.TemplateCacheImpl{}
}

func (inst* p90e018b46a_itemplates_TemplateCacheImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p90e018b46.TemplateCacheImpl)
	nop(ie, com)

	
    com.Next = inst.getNext(ie)


    return nil
}


func (inst*p90e018b46a_itemplates_TemplateCacheImpl) getNext(ie application.InjectionExt)p69cb21798.TemplateLoader{
    return ie.GetComponent("#alias-69cb21798ba841067147ba8fee303bf2-TemplateLoader").(p69cb21798.TemplateLoader)
}



// type p90e018b46.TemplateLoaderImpl in package:github.com/starter-go/mails/src/lib/golang/mailelements/itemplates
//
// id:com-90e018b46a50d1bf-itemplates-TemplateLoaderImpl
// class:
// alias:alias-69cb21798ba841067147ba8fee303bf2-TemplateLoader
// scope:singleton
//
type p90e018b46a_itemplates_TemplateLoaderImpl struct {
}

func (inst* p90e018b46a_itemplates_TemplateLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-90e018b46a50d1bf-itemplates-TemplateLoaderImpl"
	r.Classes = ""
	r.Aliases = "alias-69cb21798ba841067147ba8fee303bf2-TemplateLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p90e018b46a_itemplates_TemplateLoaderImpl) new() any {
    return &p90e018b46.TemplateLoaderImpl{}
}

func (inst* p90e018b46a_itemplates_TemplateLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p90e018b46.TemplateLoaderImpl)
	nop(ie, com)

	
    com.I18 = inst.getI18(ie)
    com.AppContext = inst.getAppContext(ie)


    return nil
}


func (inst*p90e018b46a_itemplates_TemplateLoaderImpl) getI18(ie application.InjectionExt)p52d293b71.Service{
    return ie.GetComponent("#alias-52d293b7197ae4adfab092ade2e718a2-Service").(p52d293b71.Service)
}


func (inst*p90e018b46a_itemplates_TemplateLoaderImpl) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}



// type p90e018b46.TemplateManagerImpl in package:github.com/starter-go/mails/src/lib/golang/mailelements/itemplates
//
// id:com-90e018b46a50d1bf-itemplates-TemplateManagerImpl
// class:
// alias:alias-69cb21798ba841067147ba8fee303bf2-TemplateManager
// scope:singleton
//
type p90e018b46a_itemplates_TemplateManagerImpl struct {
}

func (inst* p90e018b46a_itemplates_TemplateManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-90e018b46a50d1bf-itemplates-TemplateManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-69cb21798ba841067147ba8fee303bf2-TemplateManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p90e018b46a_itemplates_TemplateManagerImpl) new() any {
    return &p90e018b46.TemplateManagerImpl{}
}

func (inst* p90e018b46a_itemplates_TemplateManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p90e018b46.TemplateManagerImpl)
	nop(ie, com)

	
    com.Next = inst.getNext(ie)


    return nil
}


func (inst*p90e018b46a_itemplates_TemplateManagerImpl) getNext(ie application.InjectionExt)p69cb21798.TemplateCache{
    return ie.GetComponent("#alias-69cb21798ba841067147ba8fee303bf2-TemplateCache").(p69cb21798.TemplateCache)
}


