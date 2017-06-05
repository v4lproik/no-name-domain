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

	PotentialCredentials []PotentialCredentials
}

func NewForm() (f *Form) {
	return &Form{"", "", "","", "", "", "", "", make(map[string]string), "", "", "", make([]PotentialCredentials, 10)}
}
