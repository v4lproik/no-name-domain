package domain

type Form struct {
	Domain string

	ScreenShot string

	UrlForm string
	UrlToSubmit string
	UsernameArg string
	PasswordArg string
	SubmitArg string
	CsrfArg string
	OtherArgWithValue map[string]string
	MethodSubmitArg string

	FaviconMD5Hash string
	FaviconPath string

	PotentialUsername string
	PotentialPassword string
}

func NewForm() (f *Form) {
	return &Form{"", "", "","", "", "", "", "", make(map[string]string), "", "", "", "", ""}
}
