package main

import (
	"os"

	"golang.org/x/term"
)

func main() {
	fileDescriptor := int(os.Stdin.Fd())

	// going to raw mode to enable input reading
	oldstate, err := term.MakeRaw(fileDescriptor)

	if err != nil {
		panic(err)
	}

	defer term.Restore(fileDescriptor, oldstate) //restores the state of terminal on the exit of the application

	buffer := make([]byte, 3) // allrow keys are 3 bytes
	game("")

	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			panic(err)
		}

		if n == 3 && buffer[0] == 27 && buffer[1] == 91 {
			switch buffer[2] {
			case 65:
				game("UP")
			case 66:
				game("DOWN")
			case 67:
				game("RIGHT")
			case 68:
				game("LEFT")
			}
		}

		if n == 1 {
			switch buffer[0] {
			case 'q':
				game("EXIT")
				return
			}
		}

	}

}
