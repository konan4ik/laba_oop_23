package main

import (
	cm "laba/command"
	ns "laba/notification_system"
	obs "laba/observer"
	p "laba/patient"
)

type Event struct {
	Name string
}

func main() {
	system := ns.Instance()
	doctor := obs.NewDoctor("Петя")
	system.AddObserver(doctor)
	patient := p.NewPatient("Вася", 1)
	command1 := cm.RegisterPatient{Service: *system, Patient: patient}
	command1.Execute()
}
