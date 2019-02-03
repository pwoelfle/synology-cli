package describe

import (
	"encoding/json"
	"fmt"

	rootcmd "github.com/pwoelfle/synology-cli/cli/cmd"
	taskcmd "github.com/pwoelfle/synology-cli/cli/cmd/downloadstation/task"
	"github.com/pwoelfle/synology-cli/pkg/synology/api/v1/downloadstation/task"
	"github.com/spf13/cobra"
)

var (
	flagFormatted *bool
)

var DescribeCmd = &cobra.Command{
	Use:   "describe [task id]",
	Short: "Describe a specific download task.",
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := task.ID(args[0])
		getInfoRequest := task.NewGetInfoRequest([]task.ID{taskID},
			task.GetInfoRequestAdditionalDetail,
			task.GetInfoRequestAdditionalFile,
			task.GetInfoRequestAdditionalPeer,
			task.GetInfoRequestAdditionalTracker,
			task.GetInfoRequestAdditionalTransfer,
		)
		taskGetInfo := &task.TaskGetInfo{}

		if err := rootcmd.SynologyClient.Call(getInfoRequest, taskGetInfo); err != nil {
			return err
		}

		var taskJSON []byte
		var err error
		if flagFormatted != nil && *flagFormatted {
			taskJSON, err = json.MarshalIndent(taskGetInfo.Tasks[0], "", "\t")
			if err != nil {
				return err
			}
		} else {
			taskJSON, err = json.Marshal(taskGetInfo.Tasks[0])
			if err != nil {
				return err
			}

		}
		fmt.Println(string(taskJSON))

		return nil
	},
}

func init() {
	flagFormatted = DescribeCmd.Flags().Bool("formatted", true, "Format the output")

	taskcmd.TaskCmd.AddCommand(DescribeCmd)
}
