package files

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"

	"github.com/mark-humane/gh-migrate-packages/internal/utils"
)

func OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	// Return the file without closing it
	return file, nil
}

func RemoveFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}

func CreateJSON(data interface{}, filename string) error {
	// Create a new file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new JSON encoder and write to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func CreateCSV(data [][]string, filename string) error {
	utils.EnsureDirExists(filename)
	// Create a new file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV writer and write to the file
	writer := bufio.NewWriter(file)
	for _, record := range data {
		_, err = writer.WriteString(strings.Join(record, ",") + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}

func ReadCSV(filename string) ([][]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader and read from the file
	reader := bufio.NewReader(file)
	var data [][]string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		data = append(data, strings.Split(strings.TrimSpace(line), ","))
	}

	return data, nil
}

func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}
