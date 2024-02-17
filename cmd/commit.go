package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewCommiter(logger *zap.Logger) *cobra.Command {
	commitCmd := &cobra.Command{
		Use:   "commit",
		Short: "Commit the changes",
		Run: func(cmd *cobra.Command, args []string) {
			execCmd := exec.Command("/bin/bash", "./script/commiter.sh")
			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr

			err := execCmd.Run()
			if err != nil {
				logger.Error("failed to run the script to commit and push the changes", zap.String("error", err.Error()))
				return
			}
			logger.Info("Commited the changes to the daily branch")
		},
	}
	return commitCmd
}
