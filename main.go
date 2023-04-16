package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	wordle "toys/wordle/src"
)

func main() {
	i := 1

	for i != 0 {
		wordle.Game()
		fmt.Print("Want To Continue Playing (Y To Continue/ N To Stop): ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = text[:1]
		if strings.ToLower(text) == "n" || strings.ToLower(text) == "no" {
			i = 0
		}
	}
}
