package console

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinimize(t *testing.T) {
	err := ConsoleMinimize()
	require.Nil(t, err)
}

func TestMaximize(t *testing.T) {
	err := ConsoleMaximize()
	require.Nil(t, err)
}

func TestHide(t *testing.T) {
	err := ConsoleHide()
	require.Nil(t, err)
}

func TestRestore(t *testing.T) {
	err := ConsoleRestore()
	require.Nil(t, err)
}
