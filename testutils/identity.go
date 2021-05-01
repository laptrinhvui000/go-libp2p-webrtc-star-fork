package testutils

import (
	"crypto/rand"
	"testing"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/stretchr/testify/require"
)

func MustCreatePrivateKey(t *testing.T) crypto.PrivKey {
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	require.NoError(t, err)
	return priv
}

func MustCreatePeerIdentity(t *testing.T, privKey crypto.PrivKey) peer.ID {
	pid, err := peer.IDFromPublicKey(privKey.GetPublic())
	require.NoError(t, err)
	return pid
}
