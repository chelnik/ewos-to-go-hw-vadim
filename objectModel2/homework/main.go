package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	readGetFile()
}

func readGetFile() ([]byte, error) {
	// Проверяем наличие аргументов командной строки
	if len(os.Args) == 2 {
		// Определение флага для файла
		filePath := flag.String("file", "", "Путь к файлу для чтения")
		flag.Parse()
		return readFile(*filePath), nil
	}
	// Проверка, был ли передан файл

	//	читаем файл из env
	filePathFromEnv := os.Getenv("FILE")
	if filePathFromEnv == "" {
		//	читаем файл из stdin
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		filePathFromStdin := scanner.Text()
		return readFile(filePathFromStdin), nil
	}
	return readFile(filePathFromEnv), nil
}

func readFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		os.Exit(1)
	}
	// Вывод данных
	fmt.Printf("Содержимое файла %s:\n%s", filePath, data)
	return data
}
