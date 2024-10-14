package proxy

import (
	"context"
	"io"
	"net"

	"github.com/cedws/iapc/iap"
	"github.com/charmbracelet/log"
)

// Listen starts a proxy server that listens on the given address and port.
func Listen(ctx context.Context, listen string, opts []iap.DialOption) {
	if err := testConn(ctx, opts); err != nil {
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

		go handleClient(ctx, opts, conn)
	}
}

func testConn(ctx context.Context, opts []iap.DialOption) error {
	tun, err := iap.Dial(ctx, opts...)
	if tun != nil {
		defer tun.Close()
	}
	return err
}

func handleClient(ctx context.Context, opts []iap.DialOption, conn net.Conn) {
	log.Debug("Client connected", "client", conn.RemoteAddr())

	tun, err := iap.Dial(ctx, opts...)
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

	log.Debug("Client disconnected", "client", conn.RemoteAddr(), "sentbytes", tun.Sent(), "recvbytes", tun.Received())
}
