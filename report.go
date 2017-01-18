package domain

type Report struct {
	Id string
	Type string

	Form Form
}

func NewReport() (r *Report) {
	return &Report{"", NewForm()}
}
