package resmodel

type NginLinkInfo struct {
	NginLink             *NginLink `json:"ngin_link"`
	Username             string    `json:"username"`
	ProfileImageURL      string    `json:"profile_image_url"`
	Description          string    `json:"description"`
	ProgrammingLanguages []string  `json:"programming_languages"`
	JobRole              string    `json:"job_role"`
}

type NginLink struct {
	NginLinkID  string        `json:"ngin_link_id"`
	SocialLinks []*SocialLink `json:"social_links"`
}

type SocialLink struct {
	PlatformName string `json:"platform_name"`
	URL          string `json:"url"`
}
