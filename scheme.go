package urlbuilder

func (u URLBuilder) WithScheme(scheme string) URLBuilder {
	u.Scheme = scheme
	return u
}

func (u URLBuilder) HTTPS() URLBuilder {
	return u.WithScheme("https")
}

func (u URLBuilder) HTTP() URLBuilder {
	return u.WithScheme("http")
}
