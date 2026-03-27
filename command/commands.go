package command

import (
	"laba/notification_system"
	"laba/patient"
)

type RegisterPatient struct {
	System  *notification_system.NotificationSystem
	Patient *patient.Patient
}

func (r RegisterPatient) Execute() {
	r.System.RegisterPatient(*r.Patient)
}

func NewRegisterPatient(system *notification_system.NotificationSystem, patient *patient.Patient) *RegisterPatient {
	return &RegisterPatient{System: system, Patient: patient}
}
