package command

import (
	ns "laba/notification_system"
	p "laba/patient"
)

type RegisterPatient struct {
	Service ns.NotificationSystem
	Patient p.Patient
}

func (r RegisterPatient) Execute() {
	r.Service.RegisterPatient(r.Patient)
}
