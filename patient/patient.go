package patient

type Status int

const (
	StatusPending Status = iota
	StatusActive
	StatusBlocked
)

type Observer interface {
	Notify(action string)
}

type Patient struct {
	Name      string
	Diagnosis Status
}
