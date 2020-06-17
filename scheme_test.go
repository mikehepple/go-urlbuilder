package urlbuilder

import (
	"testing"
)

func TestURLBuilder_HTTP(t *testing.T) {
	u := URLBuilder{}.HTTP()
	if u.Scheme != "http" {
		t.Error("u.Scheme != http")
	}
}

func TestURLBuilder_HTTPS(t *testing.T) {
	u := URLBuilder{}.HTTPS()
	if u.Scheme != "https" {
		t.Error("u.Scheme != https")
	}
}
