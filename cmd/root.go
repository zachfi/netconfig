// Copyright © 2019 Zach Leslie <xaque208@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xaque208/znet/znet"
)

var cfgFile string
var commit bool
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "netconfig",
	Short: "Configure Junos Devices",
	Long:  ``,
	Run:   netconfig,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.netconfig.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&commit, "commit", "", false, "Commit the configuration")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Increase verbosity")

	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// Search config in home directory with name ".netconfig" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".netconfig")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())
	}
}

func netconfig(cmd *cobra.Command, args []string) {
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	viper.SetDefault("netconfig.configdir", "etc/")
	viper.AutomaticEnv()

	z := znet.Znet{}
	z.LoadConfig(cfgFile)

	configDir := viper.GetString("netconfig.configdir")
	z.ConfigDir = configDir

	z.LoadData(configDir)

	l, err := z.NewLDAPClient(z.Config.Ldap)
	if err != nil {
		log.Error(err)
	}
	defer l.Close()

	hosts := z.GetNetworkHosts(l, z.Config.Ldap.BaseDN)

	for _, host := range hosts {
		log.Debugf("Opertaing on  host: %+v", host.HostName)
		z.ConfigureNetworkHost(&host, commit)
	}

}
