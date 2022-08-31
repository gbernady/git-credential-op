package op

import (
	"encoding/json"
	"errors"
	"os/exec"
)

func run(args []string, out any) error {
	cmd := exec.Command("op", args...)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	b, err := cmd.Output()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &out); err != nil {
		return err
	}
	return nil
}
