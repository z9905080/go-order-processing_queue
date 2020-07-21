package go_order_processing_queue

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	FlagCMD         = "cmd"
	FlagNSQDAddress = "nsqd_address"
)

// init func
func init() {
	viper.AutomaticEnv()

	flag.String(FlagCMD, "customer", "consumer: produce orders ,sender: write an order ")
	flag.String(FlagNSQDAddress, "127.0.0.1", "nsqd address")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}

func Start() {
	cmd := viper.GetString(FlagCMD) // retrieve value from viper

	switch cmd {
	case "consumer":
		{
			startProduceOrder()
		}
	case "sender":
		{
			startSendOrder()
		}

	}
}
