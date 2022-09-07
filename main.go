package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/gbernady/go-op"
)

// Injected at build time
var (
	buildCommit  = ""
	buildVersion = ""
)

var (
	accountFlag = flag.String("account", "", "the account to use (if more than one is available)")
	vaultFlag   = flag.String("vault", "Personal", "the vault to use; defaults to the Personal vault")
	versionFlag = flag.Bool("version", false, "prints helper and 1Password CLI versions")
)

func init() {
	// Ensure default install location for 1Password CLI is on $PATH.
	// Reference: https://developer.1password.com/docs/cli/get-started/
	os.Setenv("PATH", fmt.Sprintf("/usr/local/bin:%s", os.Getenv("PATH")))
}

func main() {
	flag.Parse()

	if *versionFlag {
		printVersion()
		return
	}

	var attr attributes
	attr.Parse(os.Stdin)

	if attr.Protocol != "https" {
		return
	}

	switch operation := flag.Arg(0); operation {
	case "get":
		res, err := get(attr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Fprintln(os.Stdout, res.String())
		}
	case "store":
		if err := store(attr); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case "erase":
		if err := erase(attr); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	default:
	}
}

func printVersion() {
	fmt.Printf("git-credential-op %s %s %s/%s %s\n", buildVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH, buildCommit)
	v, err := op.CLI{}.Version()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println("1Password CLI", v)
	}
}

func get(attr attributes) (attributes, error) {
	cli := &op.CLI{Account: *accountFlag}
	list, err := cli.ListItems(op.WithVault(*vaultFlag), op.WithCategories(op.CategoryAPICredential))
	if err != nil {
		return attr, err
	}

	for _, entry := range list {
		item, err := cli.GetItem(entry.ID, op.WithVault(*vaultFlag))
		if err != nil {
			return attr, err
		}
		if attr.Match(item) {
			if f := item.Field("username"); f != nil {
				attr.Username = f.Value
			}
			if f := item.Field("credential"); f != nil {
				attr.Password = f.Value
			}
			break
		}
	}
	return attr, nil
}

func store(attr attributes) error {
	// FIXME: use op to store api credentials in 1p
	return nil
}

func erase(attr attributes) error {
	// FIXME: use op to erase api credentials in 1p
	return nil
}
