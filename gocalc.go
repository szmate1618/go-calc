package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/szmate1618/go-calc/commandhandler"
)

func main() {

	fmt.Println(colorGreen + greeting + colorReset)

	command := ""
	err := error(nil)
	result := ""
	reader := bufio.NewReader(os.Stdin)

	for result != "exit" {

		fmt.Print(prompt)

		command, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		command = strings.TrimSpace(command)

		result, err = commandhandler.RunCommand(command)
	}

	fmt.Println(colorGreen + farewell + colorReset)
}
