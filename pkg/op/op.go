package op

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

func run(cmd []string, flags []Flag, out any) error {
	cmd = append(cmd, "--format", "json", "--iso-timestamps")
	for _, opt := range flags {
		cmd = append(cmd, opt()...)
	}
	c := exec.Command("op", cmd...)
	if errors.Is(c.Err, exec.ErrDot) {
		c.Err = nil
	}
	b, err := c.Output()
	if err != nil {
		if ee := err.(*exec.ExitError); ee != nil {
			return fmt.Errorf("%s: %s", ee, ee.Stderr)
		}
		return err
	}
	if err := json.Unmarshal(b, &out); err != nil {
		return err
	}
	return nil
}
