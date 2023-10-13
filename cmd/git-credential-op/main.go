package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/gbernady/git-credential-op/pkg/helper"
	"github.com/gbernady/git-credential-op/pkg/opcli"
)

var (
	accountFlag = flag.String("account", "", "the account to use (if more than one is available)")
	vaultFlag   = flag.String("vault", "", "the vault to use; defaults to the Personal vault")
	versionFlag = flag.Bool("version", false, "prints helper and 1Password CLI versions")
)

func init() {
	// Ensure default install location for 1Password CLI is on $PATH.
	// Reference: https://developer.1password.com/docs/cli/get-started/
	os.Setenv("PATH", fmt.Sprintf("/usr/local/bin:%s", os.Getenv("PATH")))
}

func main() {
	flag.Parse()

	op := opcli.CLI{
		Account: *accountFlag,
	}

	if *versionFlag {
		fmt.Fprintf(os.Stdout, "git-credential-op version %s\n", version())
		v, err := op.Version()
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to read op version: %v", err)
		} else {
			fmt.Fprintf(os.Stdout, "op version %s\n", v)
		}
		return
	}

	h := &helper.Helper{
		Op:    op,
		Vault: *vaultFlag,
	}
	res, err := h.Run(helper.Operation(flag.Arg(0)), helper.ParseAttributes(os.Stdin))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintln(os.Stdout, res)
	}
}

func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok || info == nil || info.Main.Version == "" {
		// binary has not been built with module support or doesn't contain a version.
		return "(unknown)"
	}
	return info.Main.Version
}
