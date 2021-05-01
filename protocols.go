package star

import (
	ma "github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multiaddr-fmt"
)

const (
	ipfsProtocolName  = "ipfs"
	P_P2P_WEBRTC_STAR = 0x01F3
)

var (
	protocolMultiaddr ma.Multiaddr
	protocol          = ma.Protocol{
		Code:  P_P2P_WEBRTC_STAR,
		Name:  "p2p-webrtc-star",
		VCode: ma.CodeToVarint(P_P2P_WEBRTC_STAR),
	}
	protoP2P_WEBRTC_STAR = ma.Protocol{
		Name: "p2p-webrtc-direct",
		Code: P_P2P_WEBRTC_STAR,
	}

	WebRTCStar = mafmt.And(
		mafmt.TCP,
		mafmt.Or(mafmt.Base(ma.P_WSS), mafmt.Base(ma.P_WS)), mafmt.Base(P_P2P_WEBRTC_STAR),
	)
)

func init() {
	var err error

	if err = ma.AddProtocol(protocol); err != nil {
		logger.Fatal(err)
	}

	protocolMultiaddr, err = ma.NewMultiaddr("/" + protocol.Name)
	if err != nil {
		logger.Fatal(err)
	}
}
