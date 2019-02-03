package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/viper"

	"github.com/pwoelfle/synology-cli/pkg/synology/api/v2/auth"

	"github.com/pwoelfle/synology-cli/pkg/synology/client"

	"github.com/spf13/cobra"
)

var (
	cfgFile              string
	synologyURL          string
	synologyAuthUsername string
	synologyAuthPassword string

	SynologyClient client.Client
)

var RootCmd = &cobra.Command{
	Use:   "synology-cli",
	Short: "Synology CLI provides access to your Synology applications.",
	Long: `Synology CLI provides access to your Synology applications.

Example:
	$> synology-cli --url <URL> -u <USERNAME> -p <PASSWORD> ds task list

Alternatively, use environment variables.
	$> export SYNOLOGY_URL=<URL>
	$> export SYNOLOGY_USERNAME=<USERNAME>
	$> export SYNOLOGY_PASSWORD=<PASSWORD>
	$> synology-cli ds task list
	`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if len(synologyURL) <= 0 {
			return fmt.Errorf("--url must be defined")
		}
		if _, err := url.ParseRequestURI(synologyURL); err != nil {
			return fmt.Errorf("given url is not valid")
		}

		if len(synologyAuthUsername) <= 0 {
			return fmt.Errorf("--username must be defined")
		}
		if len(synologyAuthPassword) <= 0 {
			return fmt.Errorf("--url must be defined")
		}

		var err error
		SynologyClient, err = client.NewClient(synologyURL)
		if err != nil {
			return err
		}

		loginRequest := auth.NewLoginRequest(synologyAuthUsername, synologyAuthPassword, "synology-cli", auth.LoginFormatCookie)
		if err = SynologyClient.Call(loginRequest, nil); err != nil {
			return err
		}

		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		logoutRequest := auth.NewLogoutRequest("synology-cli")

		if err := SynologyClient.Call(logoutRequest, nil); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	// setup environment variable support
	viper.SetEnvPrefix("synology")
	viper.BindEnv("url")
	viper.BindEnv("username")
	viper.BindEnv("password")

	RootCmd.PersistentFlags().StringVarP(&synologyURL, "url", "", viper.GetString("url"), "The url to Synology, e.g. https://synology.example.com. Alternatively, use SYNOLOGY_URL environment variable.")
	RootCmd.PersistentFlags().StringVarP(&synologyAuthUsername, "username", "u", viper.GetString("username"), "The username used for authentication. Alternatively, use SYNOLOGY_USERNAME environment variable.")
	RootCmd.PersistentFlags().StringVarP(&synologyAuthPassword, "password", "p", viper.GetString("password"), "The password used for authentication. Alternatively, use SYNOLOGY_PASSWORD environment variable.")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
