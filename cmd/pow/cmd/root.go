package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	client "github.com/textileio/powergate/api/client"
	"github.com/textileio/powergate/buildinfo"
)

var (
	fcClient *client.Client

	cmdTimeout = time.Second * 10

	rootCmd = &cobra.Command{
		Use:               "pow",
		Short:             "A client for storage and retreival of powergate data",
		Long:              `A client for storage and retreival of powergate data`,
		Version:           buildinfo.Summary(),
		DisableAutoGenTag: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlag("serverAddress", cmd.Root().PersistentFlags().Lookup("serverAddress"))
			checkErr(err)

			target := viper.GetString("serverAddress")

			fcClient, err = client.NewClient(target)
			checkErr(err)
		},
	}
)

// Execute runs the root command.
func Execute() {
	rootCmd.SetVersionTemplate("pow build info:\n{{.Version}}\n")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("serverAddress", "127.0.0.1:5002", "address of the powergate service api")
}

func initConfig() {
	viper.SetEnvPrefix("POW")
	viper.AutomaticEnv()
}
