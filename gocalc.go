package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	command := ""
	reader := bufio.NewReader(os.Stdin)

	for command != "exit" {

		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		command = strings.TrimSpace(line)

	}
	fmt.Println("Hello, World!")
}
