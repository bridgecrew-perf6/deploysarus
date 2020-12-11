package commands

import (
	"github.com/cjaewon/deploysarus/server"
	"github.com/cjaewon/deploysarus/utils/config"
	"github.com/spf13/cobra"
)

var (
	// StartCmd ...
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the Deploysarus",
		RunE:  startFn,
	}
	startCmdFlags = struct {
		configFile string
	}{}
)

func init() {
	StartCmd.Flags().StringVar(&startCmdFlags.configFile, "config", "config.yml", "Set config file path")
}

func startFn(cmd *cobra.Command, args []string) error {
	if err := config.Load(startCmdFlags.configFile); err != nil {
		return err
	}

	server.Listen()

	return nil
}
