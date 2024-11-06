package file_utility

import (
	"errors"
	"strconv"
	"fmt"
	"os"
)

func ReadFloatFromFile(file string) (float64, error) {
	
	data, err := os.ReadFile(file)
	if err != nil {
		return 0, errors.New("unable to read the file")
	}
	floatInStringVal := string(data)

	floatInFloatVal, err := strconv.ParseFloat(floatInStringVal, 64)
	if err != nil {
		return 0, errors.New("unable to convert the balance to float64") 
	}

	return floatInFloatVal, nil
}

func WriteFloatToFile(float float64, file string) {
	data := fmt.Sprintf("%f", float)
	os.WriteFile(file, []byte(data), 0644)
	// log.Println("Balance successfully recorded")
}