package list

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	humanize "github.com/dustin/go-humanize"
	rootcmd "github.com/pwoelfle/synology-cli/cli/cmd"
	taskcmd "github.com/pwoelfle/synology-cli/cli/cmd/downloadstation/task"
	"github.com/pwoelfle/synology-cli/pkg/synology/api/v1/downloadstation/task"
	"github.com/spf13/cobra"
)

type column func(t task.Task) string

var (
	columnTaskMappers = map[string]column{
		"ID":    func(t task.Task) string { return string(t.ID) },
		"TITLE": func(t task.Task) string { return t.Title },
		"TYPE":  func(t task.Task) string { return string(t.Type) },
		"SIZE": func(t task.Task) string {
			if flagSizeInBytes != nil && *flagSizeInBytes {
				return strconv.Itoa(t.Size)
			}
			return humanize.Bytes(uint64(t.Size))
		},
		"STATUS": func(t task.Task) string { return string(t.Status) },
		"URI": func(t task.Task) string {
			if t.Additional == nil {
				return ""
			}
			return string(t.Additional.Detail.URI)
		},
		"DOWNLOAD_SPEED": func(t task.Task) string {
			if t.Additional == nil || t.Status != task.StatusDownloading {
				return ""
			}
			return humanize.Bytes(uint64(t.Additional.Transfer.SpeedDownload)) + "/s"
		},
		"UPLOAD_SPEED": func(t task.Task) string {
			if t.Additional == nil || t.Status != task.StatusDownloading {
				return ""
			}
			return humanize.Bytes(uint64(t.Additional.Transfer.SpeedUpload)) + "/s"
		},
	}
)

var (
	flagOffset      = task.ListRequestDefaultOffset
	flagLimit       = task.ListRequestDefaultLimit
	flagColumns     = "ID,TITLE,SIZE,STATUS,DOWNLOAD_SPEED,UPLOAD_SPEED"
	flagSizeInBytes *bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available download tasks.",
	RunE: func(cmd *cobra.Command, args []string) error {
		listRequest := task.NewListRequest(flagOffset, flagLimit, task.ListRequestAdditionalDetail, task.ListRequestAdditionalTransfer)
		taskList := &task.TaskList{}

		if err := rootcmd.SynologyClient.Call(listRequest, taskList); err != nil {
			return err
		}

		if err := printTaskList(taskList); err != nil {
			return err
		}

		return nil
	},
}

func printTaskList(l *task.TaskList) error {
	columns := strings.Split(flagColumns, ",")
	for _, column := range columns {
		if _, ok := columnTaskMappers[column]; !ok {
			return fmt.Errorf("column %s is unknown", column)
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	fmt.Fprintln(w, strings.Join(columns, "\t"))

	for _, task := range l.Tasks {
		line := ""
		for _, column := range columns {
			if len(line) > 0 {
				line = line + "\t"
			}
			line = line + columnTaskMappers[column](task)
		}
		fmt.Fprintln(w, line)
	}
	w.Flush()

	return nil
}

func init() {
	listCmd.Flags().IntVar(&flagOffset, "offset", flagOffset, "Offset at which list request should start")
	listCmd.Flags().IntVar(&flagLimit, "limit", flagLimit, "Maximal number which list request should contain")
	listCmd.Flags().StringVar(&flagColumns, "columns", flagColumns, "Columns that should be displayed")
	flagSizeInBytes = listCmd.Flags().Bool("size-in-bytes", false, "Display size column in bytes")

	taskcmd.TaskCmd.AddCommand(listCmd)
}
