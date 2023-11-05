package proxy

import (
	"context"
	"io"
	"net"

	"github.com/cedws/iapc/iap"
	"github.com/charmbracelet/log"
)

// Start starts a proxy server that listens on the given address and port.
func Start(listen string, opts []iap.DialOption) {
	if err := testConn(opts); err != nil {
		log.Fatalf("Error testing connection: %v", err)
	}

	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Listening", "addr", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleClient(opts, conn)
	}
}

func testConn(opts []iap.DialOption) error {
	tun, err := iap.Dial(context.Background(), opts...)
	if tun != nil {
		defer tun.Close()
	}
	return err
}

func handleClient(opts []iap.DialOption, conn net.Conn) {
	log.Info("Client connected", "client", conn.RemoteAddr())

	log.Debug("Dialing IAP", "client", conn.RemoteAddr())

	tun, err := iap.Dial(context.Background(), opts...)
	if err != nil {
		log.Errorf("Error dialing IAP: %v", err)
		return
	}
	defer tun.Close()

	log.Debug("Dialed IAP", "client", conn.RemoteAddr())

	go func() {
		if _, err := io.Copy(conn, tun); err != nil {
			log.Debug(err)
		}
	}()
	if _, err := io.Copy(tun, conn); err != nil {
		log.Debug(err)
	}

	log.Info("Client disconnected", "client", conn.RemoteAddr(), "sentbytes", tun.Sent(), "recvbytes", tun.Received())
}
