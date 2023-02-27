package cmd

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/charmbracelet/log"
	"golang.org/x/oauth2/google"

	"github.com/cedws/goiap/iap"
	"github.com/spf13/cobra"
)

var (
	compress   bool
	listen     string
	project    string
	instance   string
	zone       string
	iinterface string
	port       int
)

var rootCmd = &cobra.Command{
	Use:  "goiap",
	Long: "Utility for Google Cloud's Identity-Aware Proxy",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		opts := []iap.DialOption{
			iap.WithProject(project),
			iap.WithInstance(instance),
			iap.WithZone(zone),
			iap.WithInterface(iinterface),
			iap.WithPort(fmt.Sprint(port)),
		}
		if compress {
			opts = append(opts, iap.WithCompression())
		}

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
	},
}

func handleConn(opts []iap.DialOption, conn net.Conn) {
	log.Info("Client connected", "client", conn.RemoteAddr())

	tun, err := iap.Dial(context.Background(), getToken(), opts...)
	if err != nil {
		log.Info(err)
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
	rootCmd.Flags().BoolVarP(&compress, "compress", "c", false, "Enable WebSocket compression")
	rootCmd.Flags().StringVarP(&listen, "listen", "l", "127.0.0.1:0", "Listen address and port")
	rootCmd.Flags().StringVarP(&project, "project", "p", "", "Project ID")
	rootCmd.Flags().StringVarP(&instance, "instance", "i", "", "Target instance name")
	rootCmd.Flags().StringVarP(&zone, "zone", "z", "", "Target zone name")
	rootCmd.Flags().StringVarP(&iinterface, "interface", "n", "nic0", "Target network interface")
	rootCmd.Flags().IntVarP(&port, "port", "o", 22, "Target port")

	rootCmd.MarkFlagRequired("project")
}
