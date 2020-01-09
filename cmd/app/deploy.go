package app

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/linuxsuren/docker-yaml/cmd/common"
	"github.com/linuxsuren/docker-yaml/pkg"
	"github.com/spf13/cobra"
)

// NewDeployCommand create app deploy command
func NewDeployCommand(commonOpts *common.Options) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "deploy",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return
		},
	}
	return
}
