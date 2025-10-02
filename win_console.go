//go:build windows

package console

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed bin/conwinctl.exe
var data []byte

type ConsoleOp int

const (
	ConsoleOpMinimize ConsoleOp = iota
	ConsoleOpMaximize
	ConsoleOpHide
	ConsoleOpRestore
)

func ConsoleHide() error {
	return ConsoleControl(ConsoleOpHide)
}

func ConsoleMinimize() error {
	return ConsoleControl(ConsoleOpMinimize)
}

func ConsoleMaximize() error {
	return ConsoleControl(ConsoleOpMaximize)
}

func ConsoleRestore() error {
	return ConsoleControl(ConsoleOpRestore)
}

func ConsoleControl(op ConsoleOp) error {

	dir, err := os.MkdirTemp("", "conwinctl-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	cwe := filepath.Join(dir, "conwnctl.exe")
	err = os.WriteFile(cwe, data, 0700)
	if err != nil {
		return err
	}

	var arg string
	switch op {
	case ConsoleOpMinimize:
		arg = "MIN"
	case ConsoleOpMaximize:
		arg = "MAX"
	case ConsoleOpHide:
		arg = "HIDE"
	case ConsoleOpRestore:
		arg = "RESTORE"
	default:
		return fmt.Errorf("unexpected op: %v", op)
	}

	err = exec.Command(cwe, arg).Run()
	if err != nil {
		return err
	}
	return nil
}
