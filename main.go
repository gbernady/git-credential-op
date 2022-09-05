package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gbernady/go-op"
)

var (
	account = flag.String("account", "", "the account to use (if more than one is available)")
	vault   = flag.String("vault", "Personal", "the vault to use; defaults to the Personal vault")
)

func main() {
	flag.Parse()

	var attr attributes
	attr.Parse(os.Stdin)

	if attr.Protocol != "https" {
		return
	}

	switch mode := flag.Arg(0); mode {
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

func get(attr attributes) (attributes, error) {
	cli := &op.CLI{Account: *account}
	list, err := cli.ListItems(op.WithVault(*vault), op.WithCategories(op.CategoryAPICredential))
	if err != nil {
		return attr, err
	}

	for _, entry := range list {
		item, err := cli.GetItem(entry.ID, op.WithVault(*vault))
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
