package urlbuilder

import "net/url"

func (u URLBuilder) WithQuery(key string, value string, extraValues ...string) (URLBuilder, error) {
	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return u, err
	}

	values.Set(key, value)
	for _, extraValue := range extraValues {
		values.Add(key, extraValue)
	}
	u.RawQuery = values.Encode()

	return u, nil
}

func (u URLBuilder) MustWithQuery(key, value string, extraValues ...string) URLBuilder {
	u, err := u.WithQuery(key, value, extraValues...)
	if err != nil {
		panic(err)
	}
	return u
}
