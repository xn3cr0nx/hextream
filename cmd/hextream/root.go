package main

import (
	"fmt"
	"os"

	"github.com/xn3cr0nx/hextream/cmd/hextream/bendian"
	"github.com/xn3cr0nx/hextream/cmd/hextream/lendian"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "hextream",
	Short: "Hextream provide utilities functions to deal with hex slices",
	Long: `Hextream

This tool provides a set of utilities function to deal with hexadecimal slices.
The tool is a utility borned during go development to speed up process of pasting
slices of hex bytes into the code.

With this tool you can easily convert a hex string to a slice of hex to put inside your code.

Examples:

hextream 87a157f3fd --byte --line 4
> []byte{
>    0x87, 0xa1, 0x57, 0xf3, 
>    0xfd
> }
	`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Adds subdirectories command
	rootCmd.AddCommand(lendian.LendianCmd)
	rootCmd.AddCommand(bendian.BendianCmd)
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
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".hextream" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".hextream")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
