package model

type User struct {
	ClerkID              string   `json:"clerk_id"`
	NginLinkID           string   `json:"ngin_link_id"`
	Username             string   `json:"username"`
	ProfileImageURI      string   `json:"profile_image_uri"`
	UniversityName       string   `json:"university_name"`
	ProgrammingLanguages []string `json:"programming_languages"`
}
