package reqmodel

type CreateUser struct {
	NginLinkID           string        `json:"ngin_link_id"`
	Username             string        `json:"username"`
	ProfileImageURL      string        `json:"profile_image_url"`
	Description          string        `json:"description"`
	ProgrammingLanguages []string      `json:"programming_languages"`
	JobRole              string        `json:"job_role"`
	SocialLinks          []*SocialLink `json:"social_links"`
}

type SocialLink struct {
	PlatformName string `json:"platform_name"`
	URL          string `json:"url"`
}
