//go:build !windows
// +build !windows

package utils

import (
	"os/exec"
)

func RunCmd(cmd *exec.Cmd) error {
	go cmd.Start()
	return nil
}
