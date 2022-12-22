package cmd

import (
	"github.com/rahadiangg/go-note/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add note to current path",

	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		note, _ := cmd.Flags().GetString("note")

		utils.WriteNote(title, note, viper.GetViper().GetString("path"))
	},
}
