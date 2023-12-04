package helper

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAttributesParse(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		expect Attributes
	}{
		{
			name:   "Blank",
			value:  "",
			expect: Attributes{},
		},
		{
			name:  "All",
			value: "protocol=https\nhost=foo.com\npath=bar/baz.git\nusername=qux\npassword=wat\npassword_expiry_utc=1183110060\noauth_refresh_token=watwat\nurl=https://qux:wat@foo.com/bar/baz.git\nwwwauth[]=basic realm=\"example.com\"\n",
			expect: Attributes{
				Protocol:          "https",
				Host:              "foo.com",
				Path:              "bar/baz.git",
				Username:          "qux",
				Password:          "wat",
				PasswordExpiry:    time.Unix(1183110060, 0),
				OAuthRefreshToken: "watwat",
				URL:               "https://qux:wat@foo.com/bar/baz.git",
				WWWAuth:           []string{`basic realm="example.com"`},
			},
		},
		{
			name:  "SkipEmptyAttr",
			value: "protocol=https\nhost=foo.com\npath=\n",
			expect: Attributes{
				Protocol: "https",
				Host:     "foo.com",
			},
		},
		{
			name:  "SkipEmptyLines",
			value: "protocol=https\n\n\nhost=foo.com\n",
			expect: Attributes{
				Protocol: "https",
				Host:     "foo.com",
			},
		},
		{
			name:  "SkipUnknownAttr",
			value: "foo=bar\nhost=foo.com\nbaz=qux\n",
			expect: Attributes{
				Host: "foo.com",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := ParseAttributes(bytes.NewBufferString(test.value))
			assert.Equal(t, &test.expect, a)
		})
	}
}

func TestAttributesString(t *testing.T) {
	tests := []struct {
		name   string
		value  Attributes
		expect string
	}{
		{
			name:   "Blank",
			value:  Attributes{},
			expect: "",
		},
		{
			name: "All",
			value: Attributes{
				Protocol:          "https",
				Host:              "foo.com",
				Path:              "bar/baz.git",
				Username:          "qux",
				Password:          "wat",
				PasswordExpiry:    time.Unix(1183110060, 0),
				OAuthRefreshToken: "watwat",
				URL:               "https://qux:wat@foo.com/bar/baz.git",
				WWWAuth:           []string{`basic realm="example.com"`},
			},
			expect: "protocol=https\nhost=foo.com\npath=bar/baz.git\nusername=qux\npassword=wat\npassword_expiry_utc=1183110060\noauth_refresh_token=watwat\nurl=https://qux:wat@foo.com/bar/baz.git\nwwwauth[]=basic realm=\"example.com\"\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, test.value.String())
		})
	}
}
