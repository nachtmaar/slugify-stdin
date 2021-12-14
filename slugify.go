package main

import (
	"bufio"
	"os"

	"github.com/gosimple/slug"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	stdinText := ""
	for scanner.Scan() {
		stdinText += scanner.Text()
	}
	print(slug.Make(stdinText))
}
