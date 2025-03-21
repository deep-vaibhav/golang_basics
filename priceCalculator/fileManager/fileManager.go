package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Reads lines from a given file and converts it to float64
func ReadLinesFromFile(fileName string) ([]float64, error) {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file.")
		return nil, err
	}

	// Helps in reading contents from a file
	scanner := bufio.NewScanner(file)
	var lines []float64
	for scanner.Scan() {
		floatPrice, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error in parsing price")
			return nil, err
		}
		lines = append(lines, floatPrice)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file content failed.")
		file.Close()
		return nil, err
	}

	return lines, nil
}

func WriteJsonToFile(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to json")
	}

	file.Close()
	return nil
}
