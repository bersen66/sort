package sorting

import "github.com/spf13/cobra"

type SortingConfig struct {
	numeric    bool
	reversed   bool
	unique     bool
	byColumn   bool
	columnNum  int32
	separator  string
	outputfile string
}

func getBoolFlag(cmd *cobra.Command, flag string) bool {
	res, _ := cmd.PersistentFlags().GetBool(flag)
	return res
}

func getColNum(cmd *cobra.Command) (int32, bool) {
	if col, err := cmd.PersistentFlags().GetInt32("column"); err == nil {
		return col, true
	}
	return 0, false
}

func getSep(cmd *cobra.Command) string {
	if res, err := cmd.PersistentFlags().GetString("separator"); err == nil {
		return res
	}
	return " "
}

func getOutputFile(cmd *cobra.Command) string {
	if res, err := cmd.PersistentFlags().GetString("outputfile"); err == nil {
		return res
	}
	return "sorted.txt"
}

func needSortByColumn(cmd *cobra.Command) (need bool, n int32, sep string) {
	n, need = getColNum(cmd)
	if need {
		sep = getSep(cmd)
	}
	return
}

func NewSortingConfig(cmd *cobra.Command, args []string) *SortingConfig {
	byClm, columnNum, sep := needSortByColumn(cmd)
	return &SortingConfig{
		numeric:    getBoolFlag(cmd, "numeric"),
		reversed:   getBoolFlag(cmd, "reversed"),
		unique:     getBoolFlag(cmd, "unique"),
		byColumn:   byClm,
		columnNum:  columnNum,
		separator:  sep,
		outputfile: getOutputFile(cmd),
	}
}
