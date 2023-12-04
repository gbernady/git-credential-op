package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type Operation string

const (
	Get   Operation = "get"
	Store Operation = "store"
	Erase Operation = "erase"
)

// Helper represents a git credential helper utilizing 1Password CLI to manage git credential.
type Helper struct {
	// Account specifies the 1Password account to use when executing commands.
	// Supported values are account shorthand, sign-in address, account ID, or user ID.
	Account string

	// Config specifies the 1Password CLI configuration directory to use.
	Config string

	// Path specifies the absolute path to the 1Password CLI executable.
	// When not set (default), exec.LookPath() will be utilized to find the `op` executable on $PATH.
	Path string

	// Vault specifies the 1Password vault to use.
	Vault string
}

// CLIVersion returns the 1Password CLI version or error if it cannot be read.
func (h *Helper) CLIVersion() (string, error) {
	b, err := h.exec("--version")
	if err != nil {
		return "", fmt.Errorf("op version: %w", err)
	}
	return strings.TrimSpace(string(b)), nil
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
	item, err := h.find(attr)
	if err != nil {
		return attr, fmt.Errorf("find item: %w", err)
	}
	if item != nil {
		if f := item.field("username"); f != nil {
			attr.Username = f.Value
		}
		if f := item.field("credential"); f != nil {
			attr.Password = f.Value
		}
	}
	return attr, nil
}

func (h *Helper) store(attr *Attributes) (*Attributes, error) {
	item, err := h.find(attr)
	if err != nil {
		return attr, fmt.Errorf("find item: %w", err)
	}

	var cmd []string
	if item == nil {
		title := fmt.Sprintf("%s (%s)", attr.Host, attr.Username)
		cmd = []string{"item", "create", "--category", "API Credential", "--title", title}
	} else {
		cmd = []string{"item", "edit", item.ID}
	}

	cmd = append(cmd, fmt.Sprintf("username[text]=%s", attr.Username))
	cmd = append(cmd, fmt.Sprintf("credential[concealed]=%s", attr.Password))
	cmd = append(cmd, fmt.Sprintf("hostname[text]=%s", attr.Host))
	if !attr.PasswordExpiry.IsZero() {
		cmd = append(cmd, fmt.Sprintf("expires[date]=%d", attr.PasswordExpiry.Unix()))
	}

	if _, err := h.exec(cmd...); err != nil {
		return attr, err
	}
	return attr, nil
}

func (h *Helper) erase(attr *Attributes) (*Attributes, error) {
	item, err := h.find(attr)
	if err != nil {
		return attr, fmt.Errorf("find item: %w", err)
	}
	if item != nil {
		if _, err := h.exec("item", "delete", item.ID, "--archive"); err != nil {
			return attr, err
		}
	}
	return attr, nil
}

func (h *Helper) find(attr *Attributes) (*opitem, error) {
	var list []opitem
	if err := h.execJSON(&list, "item", "list", "--categories", "API Credential"); err != nil {
		return nil, err
	}
	for _, entry := range list {
		var item *opitem
		if err := h.execJSON(&item, "item", "get", entry.ID); err != nil {
			return nil, err
		}
		if item.matches(attr) {
			return item, nil
		}
	}
	return nil, nil // not found is not an error
}

func (h *Helper) execJSON(v any, cmd ...string) error {
	cmd = append(cmd, "--format", "json", "--iso-timestamps")
	b, err := h.exec(cmd...)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &v)
}

func (h *Helper) exec(cmd ...string) ([]byte, error) {
	if h.Account != "" {
		cmd = append(cmd, fmt.Sprintf("--account=%s", h.Account))
	}
	if h.Config != "" {
		cmd = append(cmd, fmt.Sprintf("--config=%s", h.Config))
	}
	if h.Vault != "" {
		cmd = append(cmd, fmt.Sprintf("--vault=%s", h.Vault))
	}

	path := h.Path
	if path == "" {
		p, err := exec.LookPath("op")
		if err != nil && !errors.Is(err, exec.ErrDot) {
			return nil, err
		}
		path = p
	}

	op := &exec.Cmd{
		Path: path,
		Args: append([]string{path}, cmd...),
	}
	b, err := op.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("%s: %s", ee, ee.Stderr)
		}
		return nil, err
	}
	return b, err
}
