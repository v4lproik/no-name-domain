package domain

type Source int

const (
	SourceBruteforce Source = iota
	SourceFavicon
	SourceText
)

var sourceName = []string{
	SourceBruteforce:    "BRUTEFORCE",
	SourceFavicon: "FAVICON",
	SourceText:   "TEXT",
}

type PotentialCredentials struct {
	Username string
	Password string

	Source Source
}

func (s Source) String() string {
	return sourceName[s]
}