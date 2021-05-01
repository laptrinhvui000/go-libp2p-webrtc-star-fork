package transport

import (
	"testing"

	"github.com/dennis-tra/go-libp2p-webrtc-star"
	"github.com/dennis-tra/go-libp2p-webrtc-star/testutils"
	golog "github.com/ipfs/go-log"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	"github.com/libp2p/go-libp2p-peerstore/pstoremem"
	"github.com/libp2p/go-libp2p-testing/suites/transport"
	yamux "github.com/libp2p/go-libp2p-yamux"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pion/webrtc/v2"
)

const starAddress = "/dns4/localhost/tcp/9090/ws/p2p-webrtc-star"

func init() {
	golog.SetDebugLogging()
}

func TestProtocols(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestProtocols(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestBasic(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestBasic(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestCancel(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestCancel(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestPingPong(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestPingPong(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestStress1Conn1Stream1Msg(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestStress1Conn1Stream1Msg(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestStress1Conn1Stream100Msg(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestStress1Conn1Stream100Msg(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestStress1Conn100Stream100Msg(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestStress1Conn100Stream100Msg(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestStress1Conn1000Stream10Msg(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestStress1Conn1000Stream10Msg(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestStress1Conn100Stream100Msg10MB(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestStress1Conn100Stream100Msg10MB(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestStress50Conn10Stream50Msg(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestStress50Conn10Stream50Msg(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestStreamOpenStress(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestStreamOpenStress(t, starTransportA, starTransportB, mAddr, identityA)
}

func TestStreamReset(t *testing.T) {
	starTransportA, starTransportB, mAddr, identityA := testParameters(t)
	ttransport.SubtestStreamReset(t, starTransportA, starTransportB, mAddr, identityA)
}

//var Subtests = []func(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID){
//	// ttransport.SubtestProtocols,
//	// ttransport.SubtestBasic,
//	// ttransport.SubtestCancel,
//	// ttransport.SubtestPingPong,
//
//	// Stolen from the stream muxer test suite.
//	// ttransport.SubtestStress1Conn1Stream1Msg,
//	ttransport.SubtestStress1Conn1Stream100Msg,
//	ttransport.SubtestStress1Conn100Stream100Msg,
//	// ttransport.SubtestStress50Conn10Stream50Msg,
//	// ttransport.SubtestStress1Conn1000Stream10Msg,
//	// ttransport.SubtestStress1Conn100Stream100Msg10MB,
//	// ttransport.SubtestStreamOpenStress,
//	// ttransport.SubtestStreamReset,
//}
//
//func getFunctionName(i interface{}) string {
//	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
//}
//
//func TestStreamTransport(t *testing.T) {
//	ta, tb, maddr, peerA := testParameters(t)
//
//	for _, f := range Subtests {
//		t.Run(getFunctionName(f), func(t *testing.T) {
//			f(t, ta, tb, maddr, peerA)
//		})
//	}
//}

//
//func TestStreamSingle(t *testing.T) {
//	ta, tb, maddr, peerA := testParameters(t)
//
//	go func() {
//		log.Println(http.ListenAndServe("localhost:6060", nil))
//	}()
//	time.AfterFunc(time.Second*15, func() {
//		f, err := os.Create("stack-trace.txt")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println("PRITINGNNGNN")
//		err = pprof.Lookup("goroutine").WriteTo(f, 1)
//		if err != nil {
//			log.Fatal(err)
//		}
//	})
//
//	ttransport.SubtestStress(t, ta, tb, maddr, peerA, ttransport.Options{
//		ConnNum:   1,
//		StreamNum: 100,
//		MsgNum:    100,
//		MsgMax:    100,
//		MsgMin:    100,
//	})
//}

func testParameters(t *testing.T) (transport.Transport, transport.Transport, ma.Multiaddr, peer.ID) {
	starTransportA, identityA := mustCreateStarTransport(t)
	starTransportB, _ := mustCreateStarTransport(t)

	mAddr, err := ma.NewMultiaddr(starAddress)
	if err != nil {
		t.Fatal(err)
	}
	return starTransportA, starTransportB, mAddr, identityA
}

func mustCreateStarTransport(t *testing.T) (transport.Transport, peer.ID) {
	privKey := testutils.MustCreatePrivateKey(t)
	identity := testutils.MustCreatePeerIdentity(t, privKey)
	peerstore := pstoremem.NewPeerstore()
	muxer := yamux.DefaultTransport
	return star.New(identity, peerstore, muxer).
		WithSignalConfiguration(star.SignalConfiguration{
			URLPath: "/socket.io/?EIO=3&transport=websocket",
		}).
		WithWebRTCConfiguration(webrtc.Configuration{
			ICEServers: []webrtc.ICEServer{
				{
					URLs: []string{
						"stun:stun.l.google.com:19302",
						"stun:stun1.l.google.com:19302",
						"stun:stun2.l.google.com:19302",
						"stun:stun3.l.google.com:19302",
						"stun:stun4.l.google.com:19302",
					},
				},
			},
		}), identity
}
