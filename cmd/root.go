package cmd

import (
	"context"
	"io"
	"net"

	"github.com/charmbracelet/log"
	"golang.org/x/oauth2/google"

	"github.com/cedws/goiap/iap"
	"github.com/spf13/cobra"
)

var (
	compress bool
	listen   string
	project  string
	port     uint
)

var rootCmd = &cobra.Command{
	Use:  "goiap",
	Long: "Utility for Google Cloud's Identity-Aware Proxy",
}

func startProxy(opts []iap.DialOption) {
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Listening on TCP", "server", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(opts, conn)
	}
}

func handleConn(opts []iap.DialOption, conn net.Conn) {
	log.Info("Client connected", "client", conn.RemoteAddr())

	tun, err := iap.Dial(context.Background(), getToken(), opts...)
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&compress, "compress", "c", false, "Enable WebSocket compression")
	rootCmd.PersistentFlags().StringVarP(&listen, "listen", "l", "127.0.0.1:0", "Listen address and port")
	rootCmd.PersistentFlags().StringVarP(&project, "project", "p", "", "Project ID")
	rootCmd.PersistentFlags().UintVarP(&port, "port", "o", 22, "Target port")
	rootCmd.MarkFlagRequired("project")
}
