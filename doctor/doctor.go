package doctor

import (
	"fmt"
	"laba/event_datas"
	"laba/notification_system"
)

type Doctor struct {
	name   string
	system *notification_system.NotificationSystem
}

func (d *Doctor) Update(eventData event_datas.NotificationData) {
	str := fmt.Sprintf("Доктор %s увидел, что пациент %s зарегистрировался", d.name, eventData.Patient.GetName())
	fmt.Println(str)
}

func NewDoctor(name string, system *notification_system.NotificationSystem) *Doctor {
	d := &Doctor{name: name, system: system}
	system.OnNotification.Subscribe(d.Update)

	return d
}

func (d Doctor) GetName() string {
	return d.name
}

func (d *Doctor) Destroy() {
	d.system.OnNotification.Unsubscribe(d.Update)
}
