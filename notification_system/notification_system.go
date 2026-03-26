package notification_system

import (
	"fmt"
	o "laba/observer"
	p "laba/patient"
)

type NotificationSystem struct {
	Observers []o.Observer
	Patients  []p.Patient
}

var instance *NotificationSystem

func Instance() *NotificationSystem {
	if instance == nil {
		instance = &NotificationSystem{}
	}
	return instance
}

func (ns *NotificationSystem) AddObserver(obs o.Observer) {
	ns.Observers = append(ns.Observers, obs)
}

func (ns *NotificationSystem) RegisterPatient(patient p.Patient) {
	ns.Patients = append(ns.Patients, patient)
	ns.SendNotification(patient)
}

func (ns NotificationSystem) SendNotification(patient p.Patient) {
	for _, observer := range ns.Observers {
		str := fmt.Sprintf("Пациент %s поступил с диагнозом %s ", patient.GetName(), patient.GetDiagnosis())
		observer.Update(str)
	}
}

func (ns NotificationSystem) Call_a_doctor(p p.Patient) {

}
