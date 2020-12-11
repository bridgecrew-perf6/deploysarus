package commands

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "deploysarus <command> [flags]",
		Short: "Deploysarus is auto deployment tool based on github, gitlab, gogs webhook",
		Long: heredoc.Doc(`
			A auto deployment tool based on github, gitlab, bitbucket, gogs webhook
			More info at https://github.com/cjaewon/deploysarus
		`),
		Example: heredoc.Doc(`
			$ deploysarus start
			$ deploysarus start --config config.yml
		`),

		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
)

// Execute rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
