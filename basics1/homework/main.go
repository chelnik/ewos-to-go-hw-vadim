package main

import "fmt"

//func main() {
//	fmt.Println("\033[31mHello \033[0mWorld") // https://www.shellhacks.com/bash-colors/
//}

func drawCell(m map[string]string) {
	fmt.Printf("| %s %s: %s |\n|________________________________________|\n",
		m["nameEmoji"], m["nameKey"], m["nameValue"])
	fmt.Printf("| %s %s: %s |\n|________________________________________|\n",
		m["descriptionEmoji"], m["descriptionKey"], m["descriptionValue"])
	fmt.Printf("| %s %s: %s |\n|________________________________________|\n",
		m["priceEmoji"], m["priceKey"], m["priceValue"])
	fmt.Printf("| %s %s: %s |\n|________________________________________|\n",
		m["locationEmoji"], m["locationKey"], m["locationValue"])
	fmt.Printf("| %s %s: %s |\n|________________________________________|\n",
		m["deliveryEmoji"], m["deliveryKey"], m["deliveryValue"])
}

func main() {
	cell := map[string]string{
		"typeOfDraw":       "-",
		"nameKey":          "Название",
		"nameValue":        "Станок",
		"nameEmoji":        "\U0001F4AC",
		"descriptionKey":   "Описание",
		"descriptionValue": "Станок для дерева",
		"descriptionEmoji": "\U0001F4D4",
		"priceKey":         "Цена",
		"priceValue":       "100$",
		"priceEmoji":       "\U0001F4B5",
		"locationKey":      "Локация",
		"locationValue":    "Казань",
		"locationEmoji":    "\U0001F4CD",
		"deliveryKey":      "Доставка",
		"deliveryValue":    "Имеется",
		"deliveryEmoji":    "\U0001F4E6",
		"color":            "\033[31m",
	}
	drawCell(cell)
}
