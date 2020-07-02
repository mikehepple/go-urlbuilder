package urlbuilder_test

import (
	urlbuilder "github.com/mikehepple/go-urlbuilder"
	"testing"
)

func TestURLBuilder(t *testing.T) {

	baseURL := urlbuilder.New().HTTP().WithHost("example.org")

	tests := []struct {
		name        string
		expectedURL string
		builtURL    urlbuilder.URLBuilder
	}{
		{
			"basic example",
			"http://example.org/api/v1/comments",
			baseURL.WithStaticPath("/api/v1/comments"),
		},
		{
			"example with parameterized query",
			"http://example.org/api/v1/comments/123",
			baseURL.MustWithPathWithParameters("/api/v1/comments/?", "123"),
		},
		{
			"example with named parameterized query",
			"http://example.org/api/v1/comments/123",
			baseURL.MustPathWithNamedParameters("/api/v1/comments/:commentID",
				map[string]string{"commentID": "123"}),
		},
		{
			"example with all url features",
			"https://admin:hunter2@example.org:42/api/v1/comments/foobar?page=1&page=2&page=3#barfoo",
			urlbuilder.New().
				HTTPS().WithHostAndPort("example.org", 42).
				WithUsernamePassword("admin", "hunter2").
				MustPathWithNamedParameters("/api/v1/comments/:search", map[string]string{"search": "foobar"}).
				MustWithQuery("page", "1", "2", "3").
				WithFragment("barfoo"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.builtURL.String() != tt.expectedURL {
				t.Errorf("%s != %s", tt.builtURL, tt.expectedURL)
			}
		})
	}

	expectedURL := "https://admin:hunter2@example.org:42/api/v1/comments/foobar?page=1&page=2&page=3#barfoo"
	builtURL := urlbuilder.New().
		HTTPS().WithHostAndPort("example.org", 42).
		WithUsernamePassword("admin", "hunter2").
		MustPathWithNamedParameters("/api/v1/comments/:search", map[string]string{"search": "foobar"}).
		MustWithQuery("page", "1", "2", "3").
		WithFragment("barfoo").
		String()

	if builtURL != expectedURL {
		t.Errorf("%s != %s", builtURL, expectedURL)
	}

}
