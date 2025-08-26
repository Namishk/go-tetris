package main

import (
	"os"
	"time"

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

	game := NewGame()

	quit := make(chan bool)

	go func() {

		for {
			n, err := os.Stdin.Read(buffer)
			if err != nil {
				panic(err)
			}

			if n == 3 && buffer[0] == 27 && buffer[1] == 91 {
				switch buffer[2] {
				case 65:
					game.rotateTile()
				case 66:
					game.goDown()
				case 67:
					game.MoveRight()
				case 68:
					game.moveLeft()
				}
			}

			if n == 1 {
				switch buffer[0] {
				case 'q':
					game.exit()
					quit <- true
					return
				}
			}

		}
	}()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			game.renderGame()
		case <-quit:
			return

		}
	}

}
