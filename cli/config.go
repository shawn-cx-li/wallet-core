package main

import (
	"flag"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	defaultConfigFile = "./config/config.yaml"
)

type Config struct {
	Wallet WalletConfig
	Family string
}

type WalletConfig struct {
	Mnemonic string
	Path     string
}

var v *viper.Viper

func bindFlags() (err error) {
	flag.String("conf", "", "Abosolute wallet config file location")
	flag.String("mnemonic", "", "Mnemonic of the HD Wallet")
	flag.String("path", "", "Path of the private key")
	flag.String("family", "btc-bip44", "Wallet to create")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	return v.BindPFlags(pflag.CommandLine)
}

func bindConfigFile(configFile string) (err error) {
	if configFile == "" {
		configFile, err = filepath.Abs(defaultConfigFile)
		if err != nil {
			return
		}
	}

	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)

	return
}

func readConfig() (conf *Config, err error) {
	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&conf)
	if err != nil {
		return
	}

	mnemonic := v.GetString("mnemonic")
	path := v.GetString("path")
	family := v.GetString("family")

	if mnemonic != "" {
		conf.Wallet.Mnemonic = mnemonic
	}

	if path != "" {
		conf.Wallet.Path = path
	}

	if mnemonic != "" {
		conf.Family = family
	}

	return
}

func getConfig() (conf *Config, err error) {
	v = viper.New()

	err = bindFlags()
	if err != nil {
		return
	}

	err = bindConfigFile(v.GetString("conf"))
	if err != nil {
		return
	}

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	return readConfig()
}
