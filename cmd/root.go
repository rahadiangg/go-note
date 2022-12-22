package cmd

import (
	"fmt"
	"os"

	"github.com/rahadiangg/go-note/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "go-note",
	Short: "go-note is simple application to write note CLI based",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use -h to see the list of command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func init() {
	// add command
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)

	// flag command
	addCmd.Flags().String("title", "", "title your note")
	addCmd.Flags().String("note", "", "detail your note")
	addCmd.MarkFlagsRequiredTogether("title", "note")

	deleteCmd.Flags().Int("id", 0, "ID your note want to delete")
	deleteCmd.MarkFlagRequired("id")

	rootCmd.PersistentFlags().StringVar(&config.Path, "path", "/tmp/go-note.json", "Your path for save the note")
	viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))
}
