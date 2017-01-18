package domain

type Report struct {
	Type string

	Form Form
}

func NewReport() (r *Report) {
	return &Report{"", NewForm()}
}
