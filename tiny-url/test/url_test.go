package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	urlcontroller "github.com/edenriquez/system-designs/tiny-url/pkg/url"
)

var tests = []string{
	"example.com/url",
	"example.com/complex/url/123",
}

func TetsURLShortener(t *testing.T) {
	for _, url := range tests {
		assert.Equal(t, urlcontroller.URLShortener(url), url+"123abc")
	}
}
