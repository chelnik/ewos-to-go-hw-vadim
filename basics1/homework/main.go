package main

import (
	"fmt"
	"strings"
)

const numberOfSeparator = 40

func printRow(emoji string, key string, value string, cellType string, cellColor string) {
	fmt.Printf("%s| %s %s: %s |\n|%s|\033[0m\n", cellColor, emoji, key, value, strings.Repeat(cellType, numberOfSeparator))
}

func drawCell(m map[string]string) {
	printRow(m["nameEmoji"], m["nameKey"], m["nameValue"], m["cellType"], m["cellColor"])
	printRow(m["descriptionEmoji"], m["descriptionKey"], m["descriptionValue"], m["cellType"], m["cellColor"])
	printRow(m["priceEmoji"], m["priceKey"], m["priceValue"], m["cellType"], m["cellColor"])
	printRow(m["locationEmoji"], m["locationKey"], m["locationValue"], m["cellType"], m["cellColor"])
	printRow(m["deliveryEmoji"], m["deliveryKey"], m["deliveryValue"], m["cellType"], m["cellColor"])
}

func main() {
	cell := map[string]string{
		"cellType":         ".",
		"cellColor":        "\033[33m",
		"nameEmoji":        "\U0001F4AC",
		"nameKey":          "Название",
		"nameValue":        "Станок",
		"descriptionEmoji": "\U0001F4D4",
		"descriptionKey":   "Описание",
		"descriptionValue": "Станок для дерева",
		"priceEmoji":       "\U0001F4B5",
		"priceKey":         "Цена",
		"priceValue":       "100$",
		"locationEmoji":    "\U0001F4CD",
		"locationKey":      "Локация",
		"locationValue":    "Казань",
		"deliveryEmoji":    "\U0001F4E6",
		"deliveryKey":      "Доставка",
		"deliveryValue":    "Имеется",
	}
	drawCell(cell)
}
