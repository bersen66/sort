package sorting

import (
	"fmt"
	"os"

	"github.com/bersen66/sort/pkg/fs"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run(file string, config *SortingConfig) {
	fileContent, err := fs.ReadLines(file)
	check(err)
	processor := makeChain(config)

	fileContent = processor.process(fileContent)

	fs.Flush(fileContent, config.outputfile)
}
