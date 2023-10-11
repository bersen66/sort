package cmd

import (
	"errors"
	"os"

	"github.com/bersen66/sort/pkg/sorting"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "sort [filename] [flags]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires filename!")
		}
		_, err := os.Stat(args[0])
		if os.IsNotExist(err) {
			return errors.New("File not exists!")
		}

		return nil
	},
	Short: "My implementation of sort utility",
	Long:  `Floppa - big russian cat.`,

	Run: func(cmd *cobra.Command, args []string) {
		config := sorting.NewSortingConfig(cmd, args)
		sorting.Run(args[0], config)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("numeric", "n", false, "sorting by numeric value")
	rootCmd.PersistentFlags().BoolP("reversed", "r", false, "sorting in reversed order")
	rootCmd.PersistentFlags().BoolP("unique", "u", false, "show only unique values")
	rootCmd.PersistentFlags().Int32P("column", "k", 0, "sort by column number, you have to specify separator (s flag)")
	rootCmd.PersistentFlags().StringP("separator", "s", " ", "specify column separator")
	rootCmd.PersistentFlags().StringP("out", "o", "sorted.txt", "specify output filename")
}
