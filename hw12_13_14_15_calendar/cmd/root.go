package cmd

import (
	"log"
	"os"

	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/cmd/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "./calendar --config=/path/to/config.yaml",
	Short: "Launch calendar service",
	Long:  `There is command line interface to launch calendar service.`,
	Run: func(cmd *cobra.Command, _ []string) {
		if configPath, err := cmd.Flags().GetString("config"); err == nil {
			log.Println("Launching calendar service with config: ", configPath)
			return
		}
		log.Fatal("Failed while launching calendar ...")
	},
}

func Execute() *config.Config {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(-1)
	}

	versionFlag, _ := rootCmd.Flags().GetCount("version")
	if versionFlag > 0 {
		PrintVersion()
		os.Exit(0)
	}

	configPath, _ := rootCmd.Flags().GetString("config")
	parsedConfig, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Failed while parsing config file: %s", err)
	}

	return parsedConfig
}

func init() {
	flagSet := rootCmd.Flags()

	flagSet.StringP("config", "c", "./configs/config.toml", "Path to toml config file.")
	flagSet.CountP("version", "V", "Print calendar project version.")
}
