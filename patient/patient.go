package patient

type Status int

const (
	StatusPending Status = iota
	StatusActive
	StatusBlocked
)

type Patient struct {
	name      string
	diagnosis Status
}

func NewPatient(name string, diagnosis Status) *Patient {
	return &Patient{name: name, diagnosis: diagnosis}
}

func (p Patient) GetName() string {
	return p.name
}

func (p Patient) GetDiagnosis() string {
	return p.diagnosis.String()
}