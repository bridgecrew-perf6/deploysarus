package commands

import (
	"github.com/cjaewon/deploysarus/utils/config"
	"github.com/spf13/cobra"
)

var (
	// StartCmd ...
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the Deploysarus",
		Run:   startFn,
	}
	startCmdFlags = struct {
		configFile string
	}{}
)

func init() {
	StartCmd.Flags().StringVar(&startCmdFlags.configFile, "config-file", "config.yml", "Set config file path")
}

func startFn(cmd *cobra.Command, args []string) {
	config.LoadConfigFile(startCmdFlags.configFile)
}
