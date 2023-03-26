package proxy

import (
	"context"
	"io"
	"net"

	"github.com/cedws/iapc/iap"
	"github.com/charmbracelet/log"
	"golang.org/x/oauth2/google"
)

// Start starts a proxy server that listens on the given address and port.
func Start(listen string, opts []iap.DialOption) {
	opts = append(opts, iap.WithToken(getToken()))
	if err := testConn(opts); err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}

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

	tun, err := iap.Dial(context.Background(), opts...)
	if err != nil {
		log.Error(err)
		return
	}
	defer tun.Close()
	log.Info("Established connection with proxy", "client", conn.RemoteAddr(), "sid", tun.SessionID())

	go io.Copy(conn, tun)
	io.Copy(tun, conn)

	log.Info("Client disconnected", "client", conn.RemoteAddr())
}

func getToken() string {
	credentials, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tok, err := credentials.TokenSource.Token()
	if err != nil {
		log.Fatal(err)
	}
	return tok.AccessToken
}
