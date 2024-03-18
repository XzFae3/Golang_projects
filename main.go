package main

import (
	"fmt"
	"os"
)

func main() {
	// Bonne méthode de lancement

	if len(os.Args) != 2 {
		fmt.Println("[USAGE] : go run . <filename.txt>")
		return
	} else {
		readFile(os.Args[1])
	}
}
