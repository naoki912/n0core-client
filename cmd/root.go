package cmd

import (
	"fmt"
	"os"

	"github.com/naoki912/n0core-client/n0core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var RootCmd = &cobra.Command{
	Use:   "n0core",
	Short: "n0stack/n0core client",
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("pulsar-websocket-url", "", "ws://localhost:8080", "pulsar websocket url")
	RootCmd.PersistentFlags().StringP("pulsar-topic", "", "/property/cluster/namespace/topic", "pulsar topic")
	RootCmd.PersistentFlags().StringP("pulsar-subscription", "", "subscription", "pulsar subscription name")

	viper.BindPFlag("n0core.pulsar.websocket.url", RootCmd.PersistentFlags().Lookup("pulsar-websocket-url"))
	viper.BindPFlag("n0core.pulsar.topic", RootCmd.PersistentFlags().Lookup("pulsar-topic"))
	viper.BindPFlag("n0core.pulsar.subscription", RootCmd.PersistentFlags().Lookup("pulsar-subscription"))

	viper.BindEnv("n0core.pulsar.websocket.url", "N0CORE_PULSAR_WEBSOCKET_URL")
	viper.BindEnv("n0core.pulsar.topic", "N0CORE_PULSAR_TOPIC")
	viper.BindEnv("n0core.pulsar.subscriber", "N0CORE_PULSAR_SUBSCRIBER")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".n0core")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}

func NewN0coreClient() *n0core.Client {
	return n0core.NewClient(
		viper.GetString("n0core.pulsar.websocket.url"),
		viper.GetString("n0core.pulsar.topic"),
		viper.GetString("n0core.pulsar.subscriber"),
	)
}
