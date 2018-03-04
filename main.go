// Original Source: https://github.com/rmichela/GhostCoder-iTerm/blob/master/src/main/java/com/ryanmichela/Main.java
// Reference: https://www.iterm2.com/documentation-coprocesses.html

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: ghostcoder /path/to/file.txt")
	}

	filePath := os.Args[1]
	commands, err := readCommands(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(os.Stdin)
	end := len(commands)

	for idx := 0; idx < end; {
		byte, err := reader.ReadByte()
		if err != nil {
			log.Fatalln(err)
		}

		if byte == '`' {
			fmt.Print("\b")
			simulateTyping(commands[idx])
			idx++
		}
	}
}

func simulateTyping(command string) {
	for _, b := range command {
		char := string(b)
		sleepDuration := time.Duration(80 + 3*rand.Intn(10))

		if char == " " {
			sleepDuration = time.Duration(200 + 10*rand.Intn(5))
		}

		time.Sleep(sleepDuration * time.Millisecond)
		fmt.Print(char)
	}
	fmt.Println()
}

func readCommands(filePath string) ([]string, error) {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
