package main

import (
	"fmt"
)



func main() {

	for {
		text := ""
		fmt.Print(" > ")
		fmt.Scan(&text)
		fmt.Println("echoeing.. ", text)
	}
}
