package helper

import (
	"github.com/gbernady/git-credential-op/pkg/opcli"
)

type Operation string

const (
	Get   Operation = "get"
	Store Operation = "store"
	Erase Operation = "erase"
)

// Helper represents a git credential helper utilizing 1Password CLI to manage git credential.
type Helper struct {
	// Op allows overriding the configuration of the op binary that the helper uses.
	Op opcli.CLI

	// Vault specifies the vault to use.
	Vault string
}

// Run executes the requested operation with given attributes.
//
// If an operation is not supported or not recognized, it fails silently as expected by gitcredenials.
// See https://git-scm.com/docs/gitcredentials#_custom_helpers for more details.
func (h *Helper) Run(o Operation, attr *Attributes) (*Attributes, error) {
	if attr.Protocol != "https" {
		return nil, nil
	}
	switch o {
	case Get:
		return h.get(attr)
	case Store:
		return h.store(attr)
	case Erase:
		return h.erase(attr)
	default:
		// silently ignore unknown operations
		return nil, nil
	}
}

func (h *Helper) get(attr *Attributes) (*Attributes, error) {
	list, err := h.Op.ListItems(opcli.WithVault(h.Vault), opcli.WithCategories(opcli.CategoryAPICredential))
	if err != nil {
		return attr, err
	}
	for _, entry := range list {
		item, err := h.Op.GetItem(entry.ID, opcli.WithVault(h.Vault))
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

func (h *Helper) store(attr *Attributes) (*Attributes, error) {
	// FIXME: use opcli to store api credentials in 1p
	return nil, nil
}

func (h *Helper) erase(attr *Attributes) (*Attributes, error) {
	// FIXME: use opcli to erase api credentials in 1p
	return nil, nil
}
