package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/szmate1618/go-calc/command"
)

func main() {

	fmt.Println(colorGreen + greeting + colorReset)

	line := ""
	err := error(nil)
	result := ""
	reader := bufio.NewReader(os.Stdin)

	for result != "exit" {

		fmt.Print(prompt)

		line, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)

		result, err = command.RunCommand(line)
	}

	fmt.Println(colorGreen + farewell + colorReset)
}
