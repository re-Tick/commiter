package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func setupLogger () (*zap.Logger, error) {
	// Logger
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func Execute() error {
	// Root command
	var rootCmd = &cobra.Command{
		Use:     "commiter",
		Short:   "Commiter CLI",
	}

	// setup the logger
	logger, err := setupLogger()
	if err != nil {
		return err
	}

	// Add the commiter command
	commiter := NewCommiter(logger)

	// add the subcommands to the root command
	rootCmd.AddCommand(commiter)

	if err := rootCmd.Execute(); err != nil {
		logger.Error("failed to start the CLI.", zap.Any("error", err.Error()))
		os.Exit(1)
		return err
	}
	return nil
}
