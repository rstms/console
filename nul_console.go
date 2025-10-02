//go:build !windows

package console

type ConsoleOp int

const (
	ConsoleOpMinimize ConsoleOp = iota
	ConsoleOpMaximize
	ConsoleOpHide
	ConsoleOpRestore
)

func ConsoleHide() error {
	return nil
}

func ConsoleMinimize() error {
	return nil
}

func ConsoleMaximize() error {
	return nil
}

func ConsoleRestore() error {
	return nil
}

func ConsoleControl(op ConsoleOp) error {
	return nil
}
