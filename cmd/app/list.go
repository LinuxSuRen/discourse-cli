package app

import (
	"fmt"
	"github.com/linuxsuren/docker-yaml/cmd/common"
	"github.com/spf13/cobra"
)

// NewListCommand create app command
func NewListCommand(commonOpts *common.Options) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "list",
		Short: "List all images of applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("not support yet")
		},
	}
	return
}