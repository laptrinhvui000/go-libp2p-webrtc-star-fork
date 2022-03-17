module github.com/dennis-tra/go-libp2p-webrtc-star

go 1.16

require (
	github.com/gorilla/websocket v1.4.2
	github.com/ipfs/go-log v1.0.5
	github.com/libp2p/go-libp2p v0.13.0
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/libp2p/go-libp2p-peerstore v0.2.7
	github.com/libp2p/go-libp2p-testing v0.4.0
	github.com/libp2p/go-libp2p-yamux v0.5.1
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multiaddr-fmt v0.1.0
	github.com/pion/datachannel v1.5.2
	github.com/pion/webrtc/v3 v3.1.25
	github.com/stretchr/testify v1.7.0
)

// Remove when https://github.com/libp2p/go-libp2p-testing/pull/32 is merged
replace github.com/libp2p/go-libp2p-testing v0.4.0 => github.com/dennis-tra/go-libp2p-testing v0.4.1-0.20210501154238-e24c08f444bb
