package lfsr

import (
	"fmt"
	"os"
)

func WriteToFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка создания файла %s: %w", filename, err)
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("ошибка записи в файл %s: %w", filename, err)
		}
	}

	fmt.Printf("данные записаны в %s\n", filename)
	return nil
}
