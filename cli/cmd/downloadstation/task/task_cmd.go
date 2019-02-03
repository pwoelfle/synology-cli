package task

import (
	"github.com/pwoelfle/synology-cli/cli/cmd/downloadstation"
	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Synology DownloadStation tasks.",
}

func init() {
	downloadstation.DownloadStationCmd.AddCommand(TaskCmd)
}
