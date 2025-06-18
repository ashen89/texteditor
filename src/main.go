package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	oldState, errorTerminalRaw := term.MakeRaw(int(os.Stdin.Fd()))
	if errorTerminalRaw != nil {
		panic(errorTerminalRaw)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	boolean := true
	for boolean {
		buf := make([]byte, 8)
		n, errorRead := os.Stdin.Read(buf[:])

		if errorRead != nil {
			continue
		}

		value := strings.TrimSpace(string(buf[:n]))

		fmt.Println("byte", buf[:n], "Value", value)

		if value == "q" {
			boolean = false
		}
	}
}
