package app

import (
	"github.com/linuxsuren/docker-yaml/cmd/common"
	"github.com/spf13/cobra"
)

// NewAppCommand create app command
func NewAppCommand(commonOpts *common.Options) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "app",
		Short: "app",
	}

	cmd.AddCommand(NewDeployCommand(commonOpts), NewListCommand(commonOpts))
	return
}
