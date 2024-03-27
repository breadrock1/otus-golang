package cmd

import (
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/config"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "./calendar --config=/path/to/config.yaml",
	Short: "Launch calendar service",
	Long: `
		There is command line interface to launch calendar service. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
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
		os.Exit(1)
	}

	parsedConfig := config.NewConfig()
	return parsedConfig
}

func parseConfig() {

}

func init() {
	flagSet := rootCmd.Flags()

	flagSet.StringP("config", "c", "./config.yaml", "Path to yaml config file.")
}
