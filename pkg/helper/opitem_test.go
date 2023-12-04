package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpitemMatches(t *testing.T) {
	tests := []struct {
		name  string
		item  opitem
		attr  Attributes
		match bool
	}{
		{
			name: "HostMatch",
			item: opitem{
				Fields: []opfield{
					{ID: "username", Label: "username", Value: "qux"},
					{ID: "credential", Label: "credential", Value: "wat"},
					{ID: "hostname", Label: "hostname", Value: "foo.com"},
				},
			},
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			match: true,
		},
		{
			name: "HostMissing",
			item: opitem{
				Fields: []opfield{
					{ID: "username", Label: "username", Value: "qux"},
					{ID: "credential", Label: "credential", Value: "wat"},
				},
			},
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			match: false,
		},
		{
			name: "HostMismatch",
			item: opitem{
				Fields: []opfield{
					{ID: "username", Label: "username", Value: "qux"},
					{ID: "credential", Label: "credential", Value: "wat"},
					{ID: "hostname", Label: "hostname", Value: "bar.com"},
				},
			},
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
			},
			match: false,
		},
		{
			name: "PathMatch",
			item: opitem{
				Fields: []opfield{
					{ID: "username", Label: "username", Value: "qux"},
					{ID: "credential", Label: "credential", Value: "wat"},
					{ID: "hostname", Label: "hostname", Value: "foo.com"},
					{ID: "7kdfaup5", Label: "path", Value: "bar/baz.git"},
				},
			},
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Path:     "bar/baz.git",
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
			item: opitem{
				Fields: []opfield{
					{ID: "username", Label: "username", Value: "qux"},
					{ID: "credential", Label: "credential", Value: "wat"},
					{ID: "hostname", Label: "hostname", Value: "foo.com"},
					{ID: "7kdfaup5", Label: "path", Value: "bar/qux.git"},
				},
			},
			match: false,
		},
		{
			name: "UserMatch",
			item: opitem{
				Fields: []opfield{
					{ID: "username", Label: "username", Value: "qux"},
					{ID: "credential", Label: "credential", Value: "wat"},
					{ID: "hostname", Label: "hostname", Value: "foo.com"},
				},
			},
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Username: "qux",
			},
			match: true,
		},
		{
			name: "UserMismatch",
			attr: Attributes{
				Protocol: "https",
				Host:     "foo.com",
				Username: "wat",
			},
			item: opitem{
				Fields: []opfield{
					{ID: "username", Label: "username", Value: "qux"},
					{ID: "credential", Label: "credential", Value: "wat"},
					{ID: "hostname", Label: "hostname", Value: "foo.com"},
				},
			},
			match: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.match, test.item.matches(&test.attr))
		})
	}
}
