package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	boolean := true
	for boolean {
		buf := make([]byte, 1)
		n, _ := os.Stdin.Read(buf[:])

		value := string(buf[:n])

		fmt.Println("byte", buf, "Value", value)

		if value == "q" {
			boolean = false

			term.Restore(int(os.Stdin.Fd()), oldState)
		}
	}
}
