package commands

import (
	"github.com/dhiguero/grpc-proto-manager/internal/app/gpm/manager"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateCmdLongHelp = `
This command triggers the generation of the proto stubs for a collection of protos.
`

var generateCmdExamples = `
# Generate all the protos from the current directory.
$ gpm generate .
`

var generateCmd = &cobra.Command{
	Use:     "generate <base_path>",
	Short:   "Generate the resulting stubs for a collection of proto specs",
	Long:    generateCmdLongHelp,
	Example: generateCmdExamples,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		readConfig(args[0])

		for k, v := range viper.AllSettings() {
			log.Info().Str("key", k).Interface("value", v).Msg("viper element")
		}

		log.Info().Str("temp", viper.GetString("tempPath")).Msg("From viper")
		gpm := manager.NewManager(appConfig)
		gpm.Run(args[0])
	},
}

func init() {
	generateCmd.Flags().StringVar(&appConfig.TempPath, "tempPath", "/tmp", "Temporal file for the generation of intermediate data")
	viper.BindPFlag("tempPath", generateCmd.Flags().Lookup("tempPath"))
	rootCmd.AddCommand(generateCmd)
}

func readConfig(fromPath string) {
	viper.AddConfigPath(fromPath)
	viper.SetConfigName(".gpm") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("No config file found on given path, create a .gpm.yaml file for consistent results.")
		} else {
			log.Fatal().Err(err).Msg("unable to read configuration file")
		}
	}
	log.Info().Str("path", viper.ConfigFileUsed()).Msg("configuration loaded")
}
