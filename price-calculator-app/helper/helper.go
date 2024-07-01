package helper

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

func ReadFile(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in the file")
	}

	file.Close()
	return lines, nil
}

func ParseFloats(strings []string) ([]float64, error) {
	floats := []float64{}
	for _, str := range strings {
		flt, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float64")
		}
		floats = append(floats, flt)
	}
	return floats, nil
}

func WriteToJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create the file")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert to json")
	}

	file.Close()
	return nil
}
