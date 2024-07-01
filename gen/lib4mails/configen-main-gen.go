package lib4mails

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p1f13f945f5_implements_DispatcherManager{})
    inst.register(&p1f13f945f5_implements_DriverManagerImpl{})
    inst.register(&p1f13f945f5_implements_MainDispatcherRegistry{})
    inst.register(&p1f13f945f5_implements_MockDriver{})
    inst.register(&p1f13f945f5_implements_SMSDriver{})
    inst.register(&p1f13f945f5_implements_SMTPSenderDriver{})


    return nil
}
