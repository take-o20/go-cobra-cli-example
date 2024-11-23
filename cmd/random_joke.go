package cmd

import (
	"github.com/spf13/cobra"

	"github.com/take-o20/go-cli-example/pkg/runner"
)

var randomJokeCmd = &cobra.Command{
	Use:   "random-joke",
	Short: "get joke randomly from server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := randomJokeRunner.Run(); err != nil {
			return err
		}
		return nil
	},
}

var randomJokeRunner = runner.NewRandomJokeRuuner()
