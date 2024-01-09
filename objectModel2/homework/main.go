package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

//Операции с невалидными/отсутствующими company, created_at, id добавляем в invalid_operations

type inputDto struct {
	Company   string `json:"company"`
	Operation struct {
		Type      string `json:"type,omitempty"`
		Value     any    `json:"value,omitempty"`
		Id        any    `json:"id,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
	} `json:"operation,omitempty"`
	Type      string `json:"type,omitempty"`
	Value     any    `json:"value,omitempty"` // int float string (всегда целочисленное)
	Id        any    `json:"id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type outputDto struct {
	Company              string `json:"company"`
	ValidOperationsCount int    `json:"valid_operations_count"`
	Balance              int    `json:"balance"`
	InvalidOperations    []any  `json:"invalid_operations"`
}

func main() {
	file, err := getFile()
	if err != nil {
		log.Fatal(err)
	}
	//	записываем результат работы в файл out.json
	var in []inputDto
	err = json.Unmarshal(file, &in)
	if err != nil {
		log.Println("json.Unmarshal", err)
	}
	var i, j int
	hoofs := outputDto{
		Company:              "hoofs",
		ValidOperationsCount: 0,
		Balance:              0,
		InvalidOperations:    nil,
	}
	horns := outputDto{
		Company:              "horns",
		ValidOperationsCount: 0,
		Balance:              0,
		InvalidOperations:    nil,
	}
	for _, dto := range in {
		if dto.Company == "hoofs" {
			if dto.Value == nil && dto.Operation.Value == nil {
				hoofs.InvalidOperations = append(hoofs.InvalidOperations, dto.Id)
				continue
			}
			if dto.CreatedAt == "" && dto.Operation.CreatedAt == "" {
				hoofs.InvalidOperations = append(hoofs.InvalidOperations, dto.Id)
				continue
			}
			i++
			switch dto.Operation.Type {
			case "+":
				incomeOperationValue(dto, &hoofs)
			case "income":
				incomeOperationValue(dto, &hoofs)
			case "-":
				outcomeOperationValue(dto, &hoofs)
			case "outcome":
				outcomeOperationValue(dto, &hoofs)
			//	по дефолту заходим в обычный type
			default:
				switch dto.Type {
				case "+":
					incomeValue(dto, &hoofs)
				case "income":
					incomeValue(dto, &hoofs)
				case "-":
					outcomeValue(dto, &hoofs)
				case "outcome":
					outcomeValue(dto, &hoofs)
				default:
					i--
					hoofs.InvalidOperations = append(hoofs.InvalidOperations, dto.Id)
				}
			}
		}

		if dto.Company == "horns" {
			if dto.Value == nil && dto.Operation.Value == nil {
				horns.InvalidOperations = append(horns.InvalidOperations, dto.Id)
				continue
			}
			j++
			switch dto.Operation.Type {
			case "+":
				incomeOperationValue(dto, &horns)
			case "income":
				incomeOperationValue(dto, &horns)
			case "-":
				outcomeOperationValue(dto, &horns)
			case "outcome":
				outcomeOperationValue(dto, &horns)
			//	по дефолту заходим в обычный type
			default:
				switch dto.Type {
				case "+":
					incomeValue(dto, &horns)
				case "income":
					incomeValue(dto, &horns)
				case "-":
					outcomeValue(dto, &horns)
				case "outcome":
					outcomeValue(dto, &horns)
				default:
					j--
					horns.InvalidOperations = append(horns.InvalidOperations, dto.Id)
				}
			}
		}
	}
	hoofs.ValidOperationsCount = i
	horns.ValidOperationsCount = j
	companies := []outputDto{hoofs, horns}
	formattedOutput(companies)
	// Write data to a file
	err = writeToFile("out.json", companies)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
func writeToFile(filename string, data []outputDto) error {
	// Convert data to JSON format
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Write JSON data to the file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Data written to %s\n", filename)
	return nil
}

func formattedOutput(companies []outputDto) {
	// Format and print the sorted companies
	fmt.Println("[")
	for i, company := range companies {
		fmt.Printf("\t{\n\t\t\"company\": \"%s\",\n\t\t\"valid_operations_count\": %d,\n\t\t\"balance\": %d,\n\t\t\"invalid_operations\": [\n", company.Company, company.ValidOperationsCount, company.Balance)

		// Print each element of InvalidOperations with indentation
		for j, operation := range company.InvalidOperations {
			fmt.Printf("\t\t\t\"%v\"", operation)
			// Add a comma after each element except the last one
			if j < len(company.InvalidOperations)-1 {
				fmt.Print(",")
			}
			fmt.Println()
		}

		fmt.Print("\t\t]\n\t}")

		// Add a comma after each element except the last one
		if i < len(companies)-1 {
			fmt.Print(",")
		}
		fmt.Println()
	}
	fmt.Println("]")
}
func incomeOperationValue(inDto inputDto, outDto *outputDto) {
	//	для inDto.Operation.Value
	switch inDto.Operation.Value.(type) {
	case float64:
		v, _ := inDto.Operation.Value.(int)
		outDto.Balance += v

	case string:
		v, _ := inDto.Operation.Value.(string)
		vv, _ := strconv.Atoi(v)
		outDto.Balance += vv

	case int:
		//type assertion
		v, _ := inDto.Operation.Value.(int)
		outDto.Balance += v
	}
}
func outcomeOperationValue(inDto inputDto, outDto *outputDto) {
	//	для inDto.Operation.Value
	switch inDto.Operation.Value.(type) {
	case float64:
		v, _ := inDto.Operation.Value.(int)
		outDto.Balance -= v

	case string:
		v, _ := inDto.Operation.Value.(string)
		vv, _ := strconv.Atoi(v)
		outDto.Balance -= vv

	case int:
		//type assertion
		v, _ := inDto.Operation.Value.(int)
		outDto.Balance -= v
	}
}

func incomeValue(inDto inputDto, outDto *outputDto) {
	//	для inDto.Value
	switch inDto.Value.(type) {
	case float64:
		v, _ := inDto.Value.(int)
		outDto.Balance += v

	case string:
		v, _ := inDto.Value.(string)
		vv, _ := strconv.Atoi(v)
		outDto.Balance += vv

	case int:
		//type assertion
		v, _ := inDto.Value.(int)
		outDto.Balance += v
	}
}
func outcomeValue(inDto inputDto, outDto *outputDto) {
	//	для inDto.Value
	switch inDto.Value.(type) {
	case float64:
		v, _ := inDto.Value.(int)
		outDto.Balance -= v

	case string:
		v, _ := inDto.Value.(string)
		vv, _ := strconv.Atoi(v)
		outDto.Balance -= vv

	case int:
		//type assertion
		v, _ := inDto.Value.(int)
		outDto.Balance -= v
	}
}

// читаем файл
func getFile() ([]byte, error) {
	// Проверяем наличие аргументов командной строки
	if len(os.Args) == 3 {
		// Определение флага для файла
		filePath := flag.String("file", "", "Путь к файлу для чтения")
		flag.Parse()
		doc := readFile(*filePath)
		return doc, nil
	}
	// Проверка, был ли передан файл

	//	читаем файл из env
	filePathFromEnv := os.Getenv("FILE")
	if filePathFromEnv == "" {
		//	читаем файл из stdin
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		filePathFromStdin := scanner.Text()
		doc := readFile(filePathFromStdin)
		return doc, nil
	}
	doc := readFile(filePathFromEnv)
	return doc, nil
}

func readFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		os.Exit(1)
	}
	return data
}
