//go:build windows
// +build windows

package utils

import (
	"os/exec"
	"syscall"
)

func RunCmd(cmd *exec.Cmd) error {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go cmd.Start()
	return nil
}
