package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gbernady/git-credential-op/pkg/op"
)

var (
	account = flag.String("account", "", "the account to use")
	vault   = flag.String("vault", "", "the vault to use")
)

func main() {
	flag.Parse()

	var attr Attributes
	attr.Parse(os.Stdin)

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

func get(attr Attributes) (Attributes, error) {
	list, err := op.ListItem(
		op.WithAccount(*account),
		op.WithVault(*vault),
		op.WithCategories(op.CategoryAPICredential))
	if err != nil {
		return attr, err
	}

	for _, entry := range list {
		item, err := op.GetItem(entry.ID, op.WithAccount(*account), op.WithVault(*vault))
		if err != nil {
			return attr, err
		}
		if attr.Match(item) {
			if attr.Username == "" {
				attr.Username = item.Field("username").Value
			}
			if attr.Password == "" {
				attr.Password = item.Field("credential").Value
			}
			break
		}
	}
	return attr, nil
}

func store(attr Attributes) error {
	// FIXME: use op to store api credentials in 1p
	return nil
}

func erase(attr Attributes) error {
	// FIXME: use op to erase api credentials in 1p
	return nil
}
