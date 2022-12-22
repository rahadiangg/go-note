package cmd

import (
	"github.com/rahadiangg/go-note/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "ini configuration to current path",
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateFile(viper.GetViper().GetString("path"))
	},
}
