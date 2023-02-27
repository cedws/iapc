package cmd

import (
	"fmt"

	"github.com/cedws/goiap/iap"
	"github.com/spf13/cobra"
)

var (
	zone       string
	iinterface string
)

var instanceCmd = &cobra.Command{
	Use:  "instance",
	Long: "Create a tunnel to a remote Compute Engine instance",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		opts := []iap.DialOption{
			iap.WithProject(project),
			iap.WithInstance(args[0], zone, iinterface),
			iap.WithPort(fmt.Sprint(port)),
		}
		if compress {
			opts = append(opts, iap.WithCompression())
		}

		startProxy(opts)
	},
}

func init() {
	instanceCmd.Flags().StringVarP(&zone, "zone", "z", "", "Target zone name")
	instanceCmd.Flags().StringVarP(&iinterface, "interface", "n", "nic0", "Target network interface")
	instanceCmd.MarkFlagRequired("zone")

	rootCmd.AddCommand(instanceCmd)
}
