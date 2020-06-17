package urlbuilder

import "fmt"

func (u URLBuilder) WithHost(host string) URLBuilder {
	u.Host = host
	return u
}

func (u URLBuilder) WithHostAndPort(host string, port int) URLBuilder {
	u.Host = fmt.Sprintf("%s:%d", host, port)
	return u
}
