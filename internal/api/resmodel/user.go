package resmodel

type User struct {
	ClerkID              string    `json:"clerk_id"`
	NginLink             *NginLink `json:"ngin_link"`
	Username             string    `json:"username"`
	ProfileImageURL      string    `json:"profile_image_url"`
	Description          string    `json:"description"`
	ProgrammingLanguages []string  `json:"programming_languages"`
	JobRole              string    `json:"job_role"`
}
