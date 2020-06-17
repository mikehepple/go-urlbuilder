package urlbuilder

import (
	"net/url"
)

type URLBuilder url.URL

func New() URLBuilder {
	return URLBuilder{}
}

func (u URLBuilder) URL() url.URL {
	return url.URL(u)
}

func (u URLBuilder) String() string {
	theUrl := u.URL()
	return theUrl.String()
}

func (u URLBuilder) WithFragment(fragment string) URLBuilder {
	u.Fragment = fragment
	return u
}

func (u URLBuilder) WithUsername(username string) URLBuilder {
	u.User = url.User(username)
	return u
}

func (u URLBuilder) WithUsernamePassword(username, password string) URLBuilder {
	u.User = url.UserPassword(username, password)
	return u
}
