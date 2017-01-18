package domain

type Report struct {
	Id int
	Type string

	Form Form
}

func NewReport() (r *Report) {
	return &Report{"", NewForm()}
}
