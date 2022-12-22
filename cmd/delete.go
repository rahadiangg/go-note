package cmd

import (
	"github.com/rahadiangg/go-note/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete your note based on ID",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		utils.DeleteNote(id, viper.GetViper().GetString("path"))
	},
}
