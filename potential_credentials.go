package domain

type Source int

const (
	SourceBruteforce Source = iota
	SourceFavicon
	SourceText
)

type PotentialCredentials struct {
	Username string
	Password string

	Source Source
}