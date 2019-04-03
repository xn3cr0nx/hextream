// 2019 Patrick Jusic <patrick.jusic@protonmail.com>

package lendian

import (
	"fmt"

	"github.com/spf13/cobra"
)

// LendianCmd represents the lendian command
var LendianCmd = &cobra.Command{
	Use:   "lendian",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lendian called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// LendianCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// LendianCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
