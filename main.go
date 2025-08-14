package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input)

}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)

	x := strings.Fields(strings.ToLower(trimmed))
	return x
}
