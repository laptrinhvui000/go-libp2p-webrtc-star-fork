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
	protocolMultiaddr    ma.Multiaddr
	protoP2P_WEBRTC_STAR = ma.Protocol{
		Name:  "p2p-webrtc-star",
		Code:  P_P2P_WEBRTC_STAR,
		VCode: ma.CodeToVarint(P_P2P_WEBRTC_STAR),
	}

	WebRTCStar = mafmt.And(
		mafmt.TCP,
		mafmt.Or(mafmt.Base(ma.P_WSS), mafmt.Base(ma.P_WS)), mafmt.Base(P_P2P_WEBRTC_STAR),
	)
)

func init() {
	var err error

	if err = ma.AddProtocol(protoP2P_WEBRTC_STAR); err != nil {
		logger.Fatal(err)
	}

	protocolMultiaddr, err = ma.NewMultiaddr("/" + protoP2P_WEBRTC_STAR.Name)
	if err != nil {
		logger.Fatal(err)
	}
}
