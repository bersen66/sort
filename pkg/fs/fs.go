package fs

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string = make([]string, 0, 200)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			result = append(result, str)
		}
	}

	return result, nil
}

func Flush(data []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range data {
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}
