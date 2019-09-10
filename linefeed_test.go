package linefeed

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	if runtime.GOOS == "windows" {
		require.Equal(t, "\r\n", string(Get()))
	} else if runtime.GOOS == "darwin" {
		require.Equal(t, "\n", string(Get()))
	} else {
		require.Equal(t, "\n", string(Get()))
	}
}
