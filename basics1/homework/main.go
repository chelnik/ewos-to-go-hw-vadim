package main

import "fmt"

//func main() {
//	fmt.Println("\033[31mHello \033[0mWorld") // https://www.shellhacks.com/bash-colors/
//}

func drawCell(m map[string]string) {
	fmt.Printf("| \U0001F4AC %s: %s |\n|________________________________________|\n",
		m["nameKey"], m["nameValue"])
	fmt.Printf("| \U0001F4D4 %s: %s |\n|________________________________________|\n",
		m["descriptionKey"], m["descriptionValue"])
	fmt.Printf("| \U0001F4B5 %s: %s |\n|________________________________________|\n",
		m["priceKey"], m["priceValue"])
	fmt.Printf("| \U0001F4CD %s: %s |\n|________________________________________|\n",
		m["locationKey"], m["locationValue"])
	fmt.Printf("| \U0001F4E6 %s: %s |\n|________________________________________|\n",
		m["deliveryKey"], m["deliveryValue"])
}

func main() {
	cell := map[string]string{
		"nameKey":          "Название",
		"nameValue":        "Станок",
		"descriptionKey":   "Описание",
		"descriptionValue": "Станок для дерева",
		"priceKey":         "Цена",
		"priceValue":       "100$",
		"locationKey":      "Локация",
		"locationValue":    "Казань",
		"deliveryKey":      "Доставка",
		"deliveryValue":    "Имеется",
	}
	drawCell(cell)
}
