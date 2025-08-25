package main

import "fmt"

var count = 1

func game(input string) {
	fmt.Print("\x1b[2J") // clear screen
	fmt.Print("\x1b[H")  // bring cursor to start so we can redraw ui
	fmt.Printf("%d\r\n", count)
	fmt.Printf("%s\r\n", input)
	count += 1
}
