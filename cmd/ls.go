package cmd

import (
	"github.com/rahadiangg/go-note/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all note",

	Run: func(cmd *cobra.Command, args []string) {
		utils.ReadNote(viper.GetViper().GetString("path"))
	},
}
