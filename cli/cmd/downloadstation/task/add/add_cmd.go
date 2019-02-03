package add

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	rootcmd "github.com/pwoelfle/synology-cli/cli/cmd"
	taskcmd "github.com/pwoelfle/synology-cli/cli/cmd/downloadstation/task"
	"github.com/pwoelfle/synology-cli/pkg/synology/api/v1/downloadstation/task"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add one or more download tasks.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one download link")
		}

		if len(args) == 1 && args[0] == "-" {
			// use os.Stdin as argument
			return nil
		}

		for _, link := range args {
			if _, err := url.ParseRequestURI(link); err != nil {
				return err
			}
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		links := make([]string, 0)

		if len(args) == 1 && args[0] == "-" {
			reader := bufio.NewReader(os.Stdin)

			for {
				if link, err := reader.ReadString('\n'); err != nil {
					break
				} else {
					link = strings.TrimSuffix(link, "\n")
					links = append(links, link)
				}
			}
		} else {
			links = args
		}

		for _, link := range links {
			createRequest := task.NewCreateRequest(link)

			if err := rootcmd.SynologyClient.Call(createRequest, nil); err != nil {
				return err
			} else {
				fmt.Printf("Successfully added %s\n", link)
			}
		}

		return nil
	},
}

func init() {
	taskcmd.TaskCmd.AddCommand(AddCmd)
}
