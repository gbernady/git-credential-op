package helper

import (
	"bytes"
	"testing"
	"time"

	"github.com/gbernady/git-credential-op/pkg/opcli"
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

func TestAttributesMatch(t *testing.T) {
	tests := []struct {
		name  string
		attr  Attributes
		item  *opcli.Item
		match bool
	}{
		{
			name: "HostMatch",
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &opcli.Item{
				Title:    "Foo API Key",
				Category: opcli.CategoryAPICredential,
				Fields: []opcli.Field{
					{
						ID:    "username",
						Type:  opcli.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  opcli.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
					{
						ID:    "hostname",
						Type:  opcli.FieldTypeString,
						Label: "hostname",
						Value: "foo.com",
					},
				},
			},
			match: true,
		},
		{
			name: "HostMissing",
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &opcli.Item{
				Title:    "Foo API Key",
				Category: opcli.CategoryAPICredential,
				Fields: []opcli.Field{
					{
						ID:    "username",
						Type:  opcli.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  opcli.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
				},
			},
			match: false,
		},
		{
			name: "HostMismatch",
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &opcli.Item{
				Title:    "Foo API Key",
				Category: opcli.CategoryAPICredential,
				Fields: []opcli.Field{
					{
						ID:    "username",
						Type:  opcli.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  opcli.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
					{
						ID:    "hostname",
						Type:  opcli.FieldTypeString,
						Label: "hostname",
						Value: "bar.com",
					},
				},
			},
			match: false,
		},
		{
			name: "PathMatch",
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &opcli.Item{
				Title:    "Foo API Key",
				Category: opcli.CategoryAPICredential,
				Fields: []opcli.Field{
					{
						ID:    "username",
						Type:  opcli.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  opcli.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
					{
						ID:    "hostname",
						Type:  opcli.FieldTypeString,
						Label: "hostname",
						Value: "foo.com",
					},
					{
						ID:    "7kdfaup5ymst4ujtcvo5wl35cu",
						Type:  opcli.FieldTypeString,
						Label: "path",
						Value: "bar/baz.git",
					},
				},
			},
			match: true,
		},
		{
			name: "PathMismatch",
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &opcli.Item{
				Title:    "Foo API Key",
				Category: opcli.CategoryAPICredential,
				Fields: []opcli.Field{
					{
						ID:    "username",
						Type:  opcli.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  opcli.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
					{
						ID:    "hostname",
						Type:  opcli.FieldTypeString,
						Label: "hostname",
						Value: "foo.com",
					},
					{
						ID:    "7kdfaup5ymst4ujtcvo5wl35cu",
						Type:  opcli.FieldTypeString,
						Label: "path",
						Value: "bar/qux.git",
					},
				},
			},
			match: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.match, test.attr.Match(test.item))
		})
	}
}
