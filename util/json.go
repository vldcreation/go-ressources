package util

import (
	"encoding/json"
	"log"
	"os"
)

func LoadJSON(path string, v interface{}) error {
	if err := LoadFile(path, v); err != nil {
		return err
	}

	return nil
}

func LoadFile(path string, v interface{}) error {
	// open the file
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file: %v\n", err)
			return
		}
	}()

	// process the file
	if err := ProcessFile(file, v); err != nil {
		return err
	}

	return nil
}

func ProcessFile(file *os.File, v interface{}) error {
	// decode the file
	if err := json.NewDecoder(file).Decode(v); err != nil {
		return err
	}

	return nil
}
