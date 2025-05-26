package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Functions need to start with uppercase character to be exported and visible outside the package

func WriteValueToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 1000.00, errors.New("error reading file")
	}

	balanceText := string(data)
	value, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		fmt.Println("Error parsing balance:", err)
		return 1000.00, errors.New("error parsing stored value")
	}

	return value, nil
}
