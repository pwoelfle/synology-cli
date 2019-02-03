package downloadstation

import (
	"github.com/pwoelfle/synology-cli/cli/cmd"
	"github.com/spf13/cobra"
)

var DownloadStationCmd = &cobra.Command{
	Use:   "ds",
	Short: "Synology DownloadStation application.",
}

func init() {
	cmd.RootCmd.AddCommand(DownloadStationCmd)
}
