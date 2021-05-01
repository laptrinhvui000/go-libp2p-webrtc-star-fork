package testutils

import (
	"testing"

	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func MustCreateSignalAddr(t *testing.T, signalAddr string) multiaddr.Multiaddr {
	starMultiaddr, err := multiaddr.NewMultiaddr(signalAddr)
	require.NoError(t, err)
	return starMultiaddr
}
