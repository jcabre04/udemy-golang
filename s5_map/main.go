package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range(c) {
		fmt.Println("Hex code for\t", color,"\tis\t",hex)
	}
}

func main() {
	colors := map[string]string {
		"red": "#ff000",
		"green": "#00ff00",
		"blue": "#0000ff",
		"black": "#000000",
		"white": "#fffff",
	}

	printMap(colors)
}