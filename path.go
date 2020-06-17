package urlbuilder

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func (u URLBuilder) WithStaticPath(path string) URLBuilder {
	u.Path = path
	return u
}

func (u URLBuilder) WithPathWithParameters(inputPath string, args ...string) (URLBuilder, error) {

	path := inputPath
	if varCount := strings.Count(path, "?"); varCount != len(args) {
		return u, fmt.Errorf("path %s contains %d '?' but %d args provided", path, varCount, len(args))
	}

	for _, arg := range args {
		arg = url.PathEscape(arg)
		path = strings.Replace(path, "?", arg, 1)
	}

	u.Path = path

	return u, nil
}

func (u URLBuilder) MustWithPathWithParameters(path string, args ...string) URLBuilder {
	u, err := u.WithPathWithParameters(path, args...)
	if err != nil {
		panic(err)
	}
	return u
}

var keyMatcher = regexp.MustCompile(`[\\b/]?(:\w+)[\\b/]?\b`)

func (u URLBuilder) WithPathWithNamedParameters(inputPath string, params map[string]string) (URLBuilder, error) {

	if params == nil {
		params = map[string]string{}
	}

	path := inputPath

	for key, val := range params {
		paramRegex, err := regexp.Compile(fmt.Sprintf(`[\\b/]?(:%s)[\\b/]?\b`, key))
		if err != nil {
			return u, err
		}

		if !paramRegex.MatchString(path) {
			return u, fmt.Errorf(":%s not found in path %s", key, inputPath)
		}

		val = url.PathEscape(val)
		path, err = replaceAllSubmatches(path, paramRegex, val)
		if err != nil {
			return u, err
		}

		path = paramRegex.ReplaceAllLiteralString(path, val)
	}

	if extraKeys := keyMatcher.FindAllStringSubmatch(path, -1); extraKeys != nil {
		extraKeyNames := []string{}
		for _, key := range extraKeys {
			extraKeyNames = append(extraKeyNames, key[1])
		}
		return u, fmt.Errorf("path parameters in %s not supplied %s", inputPath, strings.Join(extraKeyNames, ", "))
	}

	u.Path = path
	return u, nil
}

func replaceAllSubmatches(input string, regex *regexp.Regexp, val string) (string, error) {
	for {
		submatches := regex.FindAllStringSubmatchIndex(input, -1)
		if submatches == nil {
			return input, nil
		}
		input = input[0:submatches[0][2]] + val + input[submatches[0][3]:]

	}
}

func (u URLBuilder) MustPathWithNamedParameters(path string, params map[string]string) URLBuilder {
	u, err := u.WithPathWithNamedParameters(path, params)
	if err != nil {
		panic(err)
	}
	return u
}
