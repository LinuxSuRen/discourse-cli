package cmd

import (
	"github.com/linuxsuren/docker-yaml/cmd/app"
	"github.com/linuxsuren/docker-yaml/cmd/common"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command {
	Use:   "dcli",
	Short: "docker cli",
}

// GetRootCmd only for test purpose
func GetRootCmd() *cobra.Command {
	return rootCmd
}

var commonOptions = common.Options{}

// GetCommonOptions only for test purpose
func GetCommonOptions() *common.Options {
	return &commonOptions
}

// register all commands here
func init()  {
	rootCmd.AddCommand(app.NewAppCommand(&commonOptions))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrln(err)
		os.Exit(1)
	}
}
