package model

type User struct {
	ClerkID              string
	NginLink             *NginLink
	Username             string
	ProfileImageURL      string
	Description          string
	ProgrammingLanguages []string
	JobRole              string
}
