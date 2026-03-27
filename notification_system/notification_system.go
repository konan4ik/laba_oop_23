package notification_system

import (
	"fmt"
	interfaces "laba/Interfaces"
	"laba/event_datas"
	p "laba/patient"
	"laba/utilities"
	"sync"
)

type NotificationSystem struct {
	OnNotification utilities.Event[event_datas.NotificationData]
	Patients       []p.Patient
}

var (
	once     sync.Once
	instance *NotificationSystem
)

func GetInstance() *NotificationSystem {
	once.Do(func() {
		instance = &NotificationSystem{}
	})
	return instance
}

func (ns *NotificationSystem) RegisterPatient(patient p.Patient) {
	ns.Patients = append(ns.Patients, patient)
	ns.OnNotification.Emit(event_datas.NotificationData{Patient: patient})
}

func (ns *NotificationSystem) CallDoctor(patient p.Patient, doctor interfaces.Doctor_data) {
	str := fmt.Sprintf("Доктор %s вызван к пациенту %s", doctor.GetName(), patient.GetName())
	fmt.Println(str)
}

func (ns *NotificationSystem) SendNotification() {
	// ???
}

func (ns *NotificationSystem) InvokeCommand(c interfaces.Command) {
	c.Execute()
}
