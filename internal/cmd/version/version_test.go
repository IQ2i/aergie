package version

import (
	"testing"
)

func TestFormat(t *testing.T) {
	expects := "ae version 1.0.0\nhttps://github.com/IQ2i/aergie/releases/tag/v1.0.0\n"
	if got := Format("1.0.0"); got != expects {
		t.Errorf("versionFormat() = %q, wants %q", got, expects)
	}
}

func TestChangelogURL(t *testing.T) {
	tag := "1.0.0"
	url := "https://github.com/IQ2i/aergie/releases/tag/v1.0.0"
	result := changelogURL(tag)
	if result != url {
		t.Errorf("expected %s to create url %s but got %s", tag, url, result)
	}

	tag = "DEV"
	url = "https://github.com/IQ2i/aergie/releases/latest"
	result = changelogURL(tag)
	if result != url {
		t.Errorf("expected %s to create url %s but got %s", tag, url, result)
	}
}
