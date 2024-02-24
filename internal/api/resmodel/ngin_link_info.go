package resmodel

type NginLinkInfo struct {
	NginLink             *NginLink
	Username             string
	ProfileImageURL      string
	Description          string
	ProgrammingLanguages []string
	JobRole              string
}

type NginLink struct {
	NginLinkID  string
	SocialLinks []*SocialLink
}

type SocialLink struct {
	PlatformName string
	URL          string
}
