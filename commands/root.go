package commands

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Version: "0.0.1",
		Use:     "deploysarus <command> [flags]",
		Short:   "Deploysarus is auto deployment tool based on github, gitlab, gogs webhooks",
		Long: heredoc.Doc(`
			A auto deployment tool based on github, gitlab, bitbucket and gogs webhooks
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

func init() {
	rootCmd.AddCommand(StartCmd)
}

// Execute rootCmd
func Execute() {
	rootCmd.Execute()
}
