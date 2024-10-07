package model

type PostOptions struct {
	Board    string
	Name     string
	Email    string
	Spoiler  bool
	Body     string
	Embed    string
	Password string
	Sage     bool
	Files    []string
}

type Thread struct {
	PostOptions
	Subject string
}

type Reply struct {
	PostOptions
	Thread int
}
