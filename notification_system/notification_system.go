package notification_system

import (
	p "laba/patient"
)

type HospitalNotificationSystem struct {
	Observers []p.Observer
}

var instance *HospitalNotificationSystem

func Instance() *HospitalNotificationSystem {
	if instance == nil {
		instance = &HospitalNotificationSystem{}
	}
	return instance
}

func (ns HospitalNotificationSystem) Register_patient(p p.Patient) {

}

func (ns HospitalNotificationSystem) Send_notification(p p.Patient) {

}

func (ns HospitalNotificationSystem) Call_a_doctor(p p.Patient) {

}
