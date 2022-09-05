package main

import (
	"bytes"
	"testing"

	"github.com/gbernady/go-op"
	"github.com/stretchr/testify/assert"
)

func TestAttributesParse(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		expect attributes
	}{
		{
			name:   "Blank",
			value:  "",
			expect: attributes{},
		},
		{
			name:  "All",
			value: "protocol=https\nhost=foo.com\npath=bar/baz.git\nusername=qux\npassword=wat\nurl=https://qux:wat@foo.com/bar/baz.git\n",
			expect: attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
				Username: "qux",
				Password: "wat",
				URL:      "https://qux:wat@foo.com/bar/baz.git",
			},
		},
		{
			name:  "SkipEmptyAttr",
			value: "protocol=https\nhost=foo.com\npath=\n",
			expect: attributes{
				Protocol: "https",
				Host:     "foo.com",
			},
		},
		{
			name:  "SkipEmptyLines",
			value: "protocol=https\n\n\nhost=foo.com\n",
			expect: attributes{
				Protocol: "https",
				Host:     "foo.com",
			},
		},
		{
			name:  "SkipUnknownAttr",
			value: "foo=bar\nhost=foo.com\nbaz=qux\n",
			expect: attributes{
				Host: "foo.com",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := &attributes{}
			a.Parse(bytes.NewBufferString(test.value))
			assert.Equal(t, &test.expect, a)
		})
	}
}

func TestAttributesString(t *testing.T) {
	tests := []struct {
		name   string
		value  attributes
		expect string
	}{
		{
			name:   "Blank",
			value:  attributes{},
			expect: "",
		},
		{
			name: "All",
			value: attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
				Username: "qux",
				Password: "wat",
				URL:      "https://qux:wat@foo.com/bar/baz.git",
			},
			expect: "protocol=https\nhost=foo.com\npath=bar/baz.git\nusername=qux\npassword=wat\nurl=https://qux:wat@foo.com/bar/baz.git\n",
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
		attr  attributes
		item  *op.Item
		match bool
	}{
		{
			name: "HostMatch",
			attr: attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &op.Item{
				Title:    "Foo API Key",
				Category: op.CategoryAPICredential,
				Fields: []op.Field{
					{
						ID:    "username",
						Type:  op.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  op.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
					{
						ID:    "hostname",
						Type:  op.FieldTypeString,
						Label: "hostname",
						Value: "foo.com",
					},
				},
			},
			match: true,
		},
		{
			name: "HostMissing",
			attr: attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &op.Item{
				Title:    "Foo API Key",
				Category: op.CategoryAPICredential,
				Fields: []op.Field{
					{
						ID:    "username",
						Type:  op.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  op.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
				},
			},
			match: false,
		},
		{
			name: "HostMismatch",
			attr: attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &op.Item{
				Title:    "Foo API Key",
				Category: op.CategoryAPICredential,
				Fields: []op.Field{
					{
						ID:    "username",
						Type:  op.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  op.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
					{
						ID:    "hostname",
						Type:  op.FieldTypeString,
						Label: "hostname",
						Value: "bar.com",
					},
				},
			},
			match: false,
		},
		{
			name: "PathMatch",
			attr: attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &op.Item{
				Title:    "Foo API Key",
				Category: op.CategoryAPICredential,
				Fields: []op.Field{
					{
						ID:    "username",
						Type:  op.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  op.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
					{
						ID:    "hostname",
						Type:  op.FieldTypeString,
						Label: "hostname",
						Value: "foo.com",
					},
					{
						ID:    "7kdfaup5ymst4ujtcvo5wl35cu",
						Type:  op.FieldTypeString,
						Label: "path",
						Value: "bar/baz.git",
					},
				},
			},
			match: true,
		},
		{
			name: "PathMismatch",
			attr: attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			item: &op.Item{
				Title:    "Foo API Key",
				Category: op.CategoryAPICredential,
				Fields: []op.Field{
					{
						ID:    "username",
						Type:  op.FieldTypeString,
						Label: "username",
						Value: "qux",
					},
					{
						ID:    "credential",
						Type:  op.FieldTypeConcealed,
						Label: "credential",
						Value: "wat",
					},
					{
						ID:    "hostname",
						Type:  op.FieldTypeString,
						Label: "hostname",
						Value: "foo.com",
					},
					{
						ID:    "7kdfaup5ymst4ujtcvo5wl35cu",
						Type:  op.FieldTypeString,
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
