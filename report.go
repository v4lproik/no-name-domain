package domain

type Report struct {
	Id string
	TypeEnum string

	Form *Form
}

func NewReport(id string, typeEnum string, form *Form) (r *Report) {
	return &Report{id, typeEnum, form}
}
