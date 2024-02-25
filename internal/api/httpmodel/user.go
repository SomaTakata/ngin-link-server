package httpmodel

type User struct {
	ClerkID              string    `json:"clerk_id"`
	NginLink             *NginLink `json:"ngin_link"`
	Username             string    `json:"username"`
	ProfileImageURL      string    `json:"profile_image_url"`
	Description          string    `json:"description"`
	ProgrammingLanguages []string  `json:"programming_languages"`
	JobRole              string    `json:"job_role"`
}

// TODO: NginLinkIDとSocialLinksをNginLinkにまとめる
type CreateUser struct {
	NginLinkID           string        `json:"ngin_link_id"`
	Username             string        `json:"username"`
	ProfileImageURL      string        `json:"profile_image_url"`
	Description          string        `json:"description"`
	ProgrammingLanguages []string      `json:"programming_languages"`
	JobRole              string        `json:"job_role"`
	SocialLinks          []*SocialLink `json:"social_links"`
}
