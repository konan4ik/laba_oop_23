package command

import (
	ns "laba/notification_system"
)

type register_patient struct {
	service ns.HospitalNotificationSystem
}

func (r register_patient) execute() {
	r.service.Register_patient()
}
