package cmd

import (
	"fmt"

	"github.com/Laisky/mockery/v2/pkg/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func NewShowConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "showconfig",
		Short: "Show the merged config",
		Long: `Print out a yaml representation of the merged config.
	This initializes viper and prints out the merged configuration between
	config files, environment variables, and CLI flags.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			config := &config.Config{}
			if err := viper.UnmarshalExact(config); err != nil {
				return errors.Wrapf(err, "failed to unmarshal config")
			}
			out, err := yaml.Marshal(config)
			if err != nil {
				return errors.Wrapf(err, "Failed to marshal yaml")
			}
			log, err := getLogger(config.LogLevel)
			if err != nil {
				panic(err)
			}
			log.Info().Msgf("Using config: %s", config.Config)

			fmt.Printf("%s", string(out))
			return nil
		},
	}
}
