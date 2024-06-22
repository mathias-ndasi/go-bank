package fileOperations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	balance := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balance), 0644)
}

func GetFloatFromFile(fileName string) (accountBalance float64, error error) {
	fileContent, error := os.ReadFile(fileName)

	if error != nil {
		return 1000, errors.New("failed to read balance from file")
	}

	balanceInString := string(fileContent)
	accountBalance, error = strconv.ParseFloat(balanceInString, 64)

	if error != nil {
		return 1000, errors.New("failed to parse balance to float from file")
	}

	return
}
