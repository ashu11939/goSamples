package main 

import "fmt"

func main() {

	var color map[string]string
	colour := make(map[string]string)
	colour["white"] = "#ffffff";
	fmt.Println(colour)
	fmt.Println(color)
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	printMap(colors)
	
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println(color , " : ", hex)
	}
}