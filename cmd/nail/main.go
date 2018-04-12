package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/PI-Victor/nail/pkg/cli"
)

var rootCmd = &cobra.Command{
	Use:   "nail - \"nail\" your microservices versions in place under one common platform",
	Short: "nail",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(
		cli.AddRepo,
		cli.CreateRelease,
		cli.ValidateRelease,
		cli.InitPlatform,
		cli.DeployRelease,
	)

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("Failed loading commands: %#v", err)
	}
}

func init() {
	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("An error occured while reading the configuration: %#v", err)
	}
	logLevel := viper.Get("log_level").(int)
	l := logrus.Level(logLevel)
	logrus.SetLevel(l)
}
