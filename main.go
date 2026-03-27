package main

import (
	cm "laba/command"
	"laba/doctor"
	ns "laba/notification_system"
	p "laba/patient"
)

type Event struct {
	Name string
}

func main() {
	system := ns.GetInstance()
	patient := p.NewPatient("Вася", 1)
	doctor1 := doctor.NewDoctor("Петя", system)
	doctor.NewDoctor("Петя2", system)
	system.InvokeCommand(cm.NewRegisterPatient(system, patient))

	doctor1.Destroy()
}
