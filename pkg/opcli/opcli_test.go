package opcli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fatih/camelcase"
	"github.com/stretchr/testify/assert"
)

func mockOp(t *testing.T) string {
	tmp := t.TempDir()

	// copy `op` mock
	b, err := os.ReadFile(filepath.Join("../../testdata", "op.sh"))
	if err != nil {
		t.Error(err)
	}
	if err := os.WriteFile(filepath.Join(tmp, "op"), b, 0744); err != nil {
		t.Error(err)
	}

	// copy test response
	tname := strings.ToLower(strings.Join(camelcase.Split(strings.ReplaceAll(t.Name(), "/", ""))[1:], "_"))
	b, err = os.ReadFile(filepath.Join("../../testdata/fixtures", fmt.Sprintf("opcli_%s", tname)))
	if err != nil {
		t.Error(err)
	}
	if err := os.WriteFile(filepath.Join(tmp, "op_response"), b, 0644); err != nil {
		t.Error(err)
	}

	return filepath.Join(tmp, "op")
}

func TestVersion(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp string
		err  string
	}{
		{
			name: "Success",
			call: func(cli *CLI) (any, error) {
				return cli.Version()
			},
			resp: `2.7.0`,
		},
		{
			name: "ExecNotFound",
			call: func(cli *CLI) (any, error) {
				cli.Path = "/foo/op"
				return cli.Version()
			},
			err: "fork/exec /foo/op: no such file or directory",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			resp, err := test.call(cli)
			if test.err == "" {
				assert.Equal(t, test.resp, resp)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}
